package notify

import (
	"fmt"
	"log"
	"strings"

	"github.com/ipron-ne/client-sdk-go/code"
	"github.com/ipron-ne/client-sdk-go/service"
	"github.com/ipron-ne/client-sdk-go/utils"
	"github.com/ipron-ne/client-sdk-go/types"
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

// set agentData(agentData)


// AddSubscriptions adds a new subscription using EventSource
func AddSubscriptions(tntId, topic string, eventCallback any, eventErrorCallback func(error), subscribePath string) error {
	client := service.GetApiClient()

	client.Lock()
	defer client.Unlock()

	params := map[string]any{
		"id":           utils.CreateUUID(),
		"clientId":     client.ClientID,
		"eventsubject": topic,
		"bcloudToken":  client.Token,
	}
	paramsString := utils.ParamsSerializer(params)

	if client.IsDebug {
		log.Printf("Try connect eventSource [%s]: %+v", topic, params)
	}

	if subscribePath == "" {
		subscribePath = strings.Split(topic, "/")[0]
	}

	fullURL := fmt.Sprintf("%s%s/%s/subscribe/%s?%s", client.BaseURL, apiName, tntId, subscribePath, paramsString)
	eventSubs, err := utils.NewEventSubscription(fullURL, "")
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic [%s]: %w", topic, err)
	}

	eventSubs.AddEventListener(string(code.Event.Handler.Register), func(e utils.Event) {
		data := utils.JSONParse(e.Data())
		if client.IsDebug {
			client.Log.Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}
		if eventCallback != nil {
			CallObjectFn(eventCallback, string(code.Event.Handler.Register), data)
		}
	})
	eventSubs.AddEventListener(string(code.Event.Handler.Registered), func(e utils.Event){
		data := utils.JSONParse(e.Data())
		if client.IsDebug {
			client.Log.Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}
		if eventCallback != nil {
			CallObjectFn(eventCallback, string(code.Event.Handler.Registered), data)
		}
	})
	eventSubs.AddEventListener(string(code.Event.Handler.Push), func(e utils.Event){
		data := utils.JSONParse(e.Data())
		if client.IsDebug {
			client.Log.Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}
		if eventCallback != nil {
			CallObjectFn(eventCallback, string(code.Event.Handler.Push), data)
		}
	})
	eventSubs.AddEventListener(string(code.Event.Handler.ProbeReq), func(e utils.Event){
		data := utils.JSONParse(e.Data())
		if client.IsDebug {
			client.Log.Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}
		if eventCallback != nil {
			CallObjectFn(eventCallback, string(code.Event.Handler.ProbeReq), data)
		}
	})
	eventSubs.AddEventListener(string(code.Event.Handler.Banishment), func(e utils.Event){
		data := utils.JSONParse(e.Data())
		if client.IsDebug {
			client.Log.Debug("%s [%s] Event %+v", e.Event(), topic, data)
		}

		clientId := client.ClientID
		oldAppID := data.Get("data").Get("oldAppId").Str()
		isBanished := (clientId == oldAppID)

		if isBanished {
			DelSubscriptions(topic)
		}

		if eventCallback != nil {
			CallObjectFn(eventCallback, string(code.Event.Handler.Banishment), data)
		}
	})
	eventSubs.OnMessage(func(e utils.Event){
		data := utils.JSONParse(e.Data())
		if client.IsDebug {
			client.Log.Debug("%s [%s]: %+v", e.Event(), topic, data)
		}
	})
	eventSubs.OnError(func(e error){
		log.Printf("Catch error eventsource [%s]: %+v", topic, e)
		DelSubscriptions(topic)

		if eventErrorCallback != nil {
			if client.IsDebug {
				client.Log.Debug("Call function eventErrorCallback")
			}
			eventErrorCallback(e)
		}
	})

	client.EventMap[topic] = eventSubs

	go eventSubs.EventLoop()

	return nil
}

func CallObjectFn(eventCallback any, id string, data types.Data) {
	switch f := eventCallback.(type) {
	case code.Function:
		f(data)
	case code.FunctionMap:
		f[id](data)
	}
}


// DelSubscriptions deletes a subscription for a given topic
func DelSubscriptions(topic string) {
	client := service.GetApiClient()

	client.Lock()
	defer client.Unlock()

	if topic != "" {
		if eventMap, exists := client.EventMap[topic]; exists {
			eventMap.EventSource.Close()
			delete(client.EventMap, topic)
			log.Printf("Closed EventSource for topic [%s]", topic)
		} else {
			log.Printf("No EventSource found for topic [%s]", topic)
		}
	} else {
		log.Println("Closing all EventSources.")
		for t, eventMap := range client.EventMap {
			eventMap.EventSource.Close()
			delete(client.EventMap, t)
		}
	}
}

// AddUserSubscriptions subscribes to user events
func AddUserSubscriptions(tntId, userId string, eventCallback func(string), eventErrorCallback func(error)) {
	topic := fmt.Sprintf("user/%s", userId)
	AddSubscriptions(tntId, topic, eventCallback, eventErrorCallback, topic)
}

// DelUserSubscriptions unsubscribes from user events
func DelUserSubscriptions(userId string) {
	topic := fmt.Sprintf("user/%s", userId)
	DelSubscriptions(topic)
}

// AddCallSubscriptions subscribes to call events
func AddCallSubscriptions(tntId, callId string, eventCallback func(string), eventErrorCallback func(error)) {
	topic := fmt.Sprintf("call/%s", callId)
	AddSubscriptions(tntId, topic, eventCallback, eventErrorCallback, topic)
}

// DelCallSubscriptions unsubscribes from call events
func DelCallSubscriptions(callId string) {
	topic := fmt.Sprintf("call/%s", callId)
	DelSubscriptions(topic)
}

// AddPhoneSubscriptions subscribes to phone events
func AddPhoneSubscriptions(tntId, phoneId string, eventCallback func(string), eventErrorCallback func(error)) {
	topic := fmt.Sprintf("phone/%s", phoneId)
	AddSubscriptions(tntId, topic, eventCallback, eventErrorCallback, topic)
}

// DelPhoneSubscriptions unsubscribes from phone events
func DelPhoneSubscriptions(phoneId string) {
	topic := fmt.Sprintf("phone/%s", phoneId)
	DelSubscriptions(topic)
}

// AddQueueSubscriptions subscribes to queue events
func AddQueueSubscriptions(tntId, queueId string, eventCallback func(string), eventErrorCallback func(error)) {
	topic := fmt.Sprintf("queue/%s", queueId)
	AddSubscriptions(tntId, topic, eventCallback, eventErrorCallback, topic)
}

// DelQueueSubscriptions unsubscribes from queue events
func DelQueueSubscriptions(queueId string) {
	topic := fmt.Sprintf("queue/%s", queueId)
	DelSubscriptions(topic)
}
