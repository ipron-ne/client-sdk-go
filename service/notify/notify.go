package notify

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/types"
	"github.com/ipron-ne/client-sdk-go/utils"
)

// Constants
const (
	apiPrefix  = "/webapi/sse"
	apiModule  = "/notify"
	apiVersion = "/v1"
)

var (
	apiName = apiPrefix + apiVersion + apiModule
	tntID   string
)

type Notify struct {
	types.Client
	eventMap   map[string]*utils.EventSubscription // topic별 수신이벤트 처리를 위한 맵
	mu         sync.Mutex
	useRegWait bool
}

func NewFromClient(client types.Client) *Notify {
	return &Notify{
		Client:     client,
		eventMap:   make(map[string]*utils.EventSubscription),
		useRegWait: true,
	}
}

func (s *Notify) SetRegWait(flag bool) {
	s.useRegWait = flag
}

func (s *Notify) IsRegWait() bool {
	return s.useRegWait
}

func (s *Notify) Lock() {
	s.mu.Lock()
}

func (s *Notify) Unlock() {
	s.mu.Unlock()
}

func (s *Notify) GetSubscriptions(topic string) *utils.EventSubscription {
	return s.eventMap[topic]
}

// AddSubscriptions adds a new subscription using EventSource
func (s *Notify) AddSubscriptions(tntId, topic string, sysEventCallback any, appEventCallback any, eventErrorCallback func(error), subscribePath string) error {
	s.Lock()
	defer s.Unlock()

	regiChan := make(chan bool, 1)
	errorMsg := ""

	params := map[string]any{
		"id":           utils.CreateUUID(),
		"clientId":     s.GetClientID(),
		"eventsubject": topic,
		"bcloudToken":  s.GetToken(),
	}
	paramsString := utils.ParamsSerializer(params)

	if s.IsDebug() {
		log.Printf("Try connect eventSource [%s]: %+v", topic, params)
	}

	if subscribePath == "" {
		subscribePath = strings.Split(topic, "/")[0]
	}

	fullURL := fmt.Sprintf("%s%s/%s/subscribe/%s?%s", s.GetBaseURL(), apiName, tntId, subscribePath, paramsString)
	eventSubs, err := utils.NewEventSubscription(fullURL, "")
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic [%s]: %w", topic, err)
	}

	eventSubs.AddEventListener(string(code.Event.Handler.Register), func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		if s.IsDebug() {
			s.GetLogger().Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}
		if sysEventCallback != nil {
			callObjectFn(sysEventCallback, string(code.Event.Handler.Register), data)
		}
	})
	// 이벤트 수신 등록 완료 이벤트
	eventSubs.AddEventListener(string(code.Event.Handler.Registered), func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		if s.IsDebug() {
			s.GetLogger().Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}
		if data.Get("result").Str() == "success" {
			eventSubs.SetRegist(true)
			regiChan <- true
		} else {
			errorMsg = data.Get("data").Get("fail").Str()
			eventSubs.SetRegist(false)
			regiChan <- false
		}
		if sysEventCallback != nil {
			callObjectFn(sysEventCallback, string(code.Event.Handler.Registered), data)
		}
	})
	eventSubs.AddEventListener(string(code.Event.Handler.Push), func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		if s.IsDebug() {
			s.GetLogger().Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}
		if appEventCallback != nil {
			callObjectFn(appEventCallback, string(code.Event.Handler.Push), data)
		}
	})
	eventSubs.AddEventListener(string(code.Event.Handler.ProbeReq), func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		if s.IsDebug() {
			s.GetLogger().Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}
		if sysEventCallback != nil {
			callObjectFn(sysEventCallback, string(code.Event.Handler.ProbeReq), data)
		}
	})
	// 강재종료 이벤트 핸들러
	eventSubs.AddEventListener(string(code.Event.Handler.Banishment), func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		if s.IsDebug() {
			s.GetLogger().Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}

		clientId := s.GetClientID()
		oldAppID := data.Get("data").Get("oldAppId").Str()
		isBanished := (clientId == oldAppID)

		if isBanished {
			s.DelSubscriptions(topic)
		}

		if sysEventCallback != nil {
			callObjectFn(sysEventCallback, string(code.Event.Handler.Banishment), data)
		}
	})
	eventSubs.OnMessage(func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		if s.IsDebug() {
			s.GetLogger().Debug("%s [%s]: %+v", e.Event(), topic, data)
		}
	})
	eventSubs.OnError(func(e error) {
		log.Printf("Catch error eventsource [%s]: %+v", topic, e)
		s.DelSubscriptions(topic)

		if eventErrorCallback != nil {
			if s.IsDebug() {
				s.GetLogger().Debug("Call function eventErrorCallback")
			}
			eventErrorCallback(e)
		}
	})

	s.eventMap[topic] = eventSubs

	go eventSubs.EventLoop()

	if s.IsRegWait() {
		select {
		case result := <-regiChan:
			if !result {
				return fmt.Errorf("AddSubscriptions: register fail. msg=%s", errorMsg)
			}
		case <-time.After(3 * time.Second):
			s.eventMap[topic] = nil
			return fmt.Errorf("AddSubscriptions: register timeout. topic=%s", topic)
		}
	}

	return nil
}

func callObjectFn(eventCallback any, id string, data types.Data) {
	switch f := eventCallback.(type) {
	case types.Function:
		f(data)
	case types.FunctionMap:
		f[id](data)
	}
}

// DelSubscriptions deletes a subscription for a given topic
func (s *Notify) DelSubscriptions(topic string) {
	s.Lock()
	defer s.Unlock()

	if topic != "" {
		if eventMap, exists := s.eventMap[topic]; exists {
			eventMap.EventSource.Close()
			delete(s.eventMap, topic)
			log.Printf("Closed EventSource for topic [%s]", topic)
		} else {
			log.Printf("No EventSource found for topic [%s]", topic)
		}
	} else {
		log.Println("Closing all EventSources.")
		for t, eventMap := range s.eventMap {
			eventMap.EventSource.Close()
			delete(s.eventMap, t)
		}
	}
}

// AddUserSubscriptions subscribes to user events
func (s *Notify) AddUserSubscriptions(tntId, userId string, eventCallback func(types.Data), eventErrorCallback func(error)) error {
	topic := fmt.Sprintf("user/%s", userId)
	return s.AddSubscriptions(tntId, topic, eventCallback, eventCallback, eventErrorCallback, topic)
}

// DelUserSubscriptions unsubscribes from user events
func (s *Notify) DelUserSubscriptions(userId string) {
	topic := fmt.Sprintf("user/%s", userId)
	s.DelSubscriptions(topic)
}

// AddCallSubscriptions subscribes to call events
func (s *Notify) AddCallSubscriptions(tntId, callId string, eventCallback func(types.Data), eventErrorCallback func(error)) error {
	topic := fmt.Sprintf("call/%s", callId)
	return s.AddSubscriptions(tntId, topic, eventCallback, eventCallback, eventErrorCallback, topic)
}

// DelCallSubscriptions unsubscribes from call events
func (s *Notify) DelCallSubscriptions(callId string) {
	topic := fmt.Sprintf("call/%s", callId)
	s.DelSubscriptions(topic)
}

// AddPhoneSubscriptions subscribes to phone events
func (s *Notify) AddPhoneSubscriptions(tntId, phoneId string, eventCallback func(types.Data), eventErrorCallback func(error)) error {
	topic := fmt.Sprintf("phone/%s", phoneId)
	return s.AddSubscriptions(tntId, topic, eventCallback, eventCallback, eventErrorCallback, topic)
}

// DelPhoneSubscriptions unsubscribes from phone events
func (s *Notify) DelPhoneSubscriptions(phoneId string) {
	topic := fmt.Sprintf("phone/%s", phoneId)
	s.DelSubscriptions(topic)
}

// AddQueueSubscriptions subscribes to queue events
func (s *Notify) AddQueueSubscriptions(tntId, queueId string, eventCallback func(types.Data), eventErrorCallback func(error)) error {
	topic := fmt.Sprintf("queue/%s", queueId)
	return s.AddSubscriptions(tntId, topic, eventCallback, eventCallback, eventErrorCallback, topic)
}

// DelQueueSubscriptions unsubscribes from queue events
func (s *Notify) DelQueueSubscriptions(queueId string) {
	topic := fmt.Sprintf("queue/%s", queueId)
	s.DelSubscriptions(topic)
}
