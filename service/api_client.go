package service

import (
	"encoding/json"
	"bytes"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/donovanhide/eventsource"

	"github.com/ipron-ne/client-sdk-go/utils"
)

const AUTH_HEADER   = "Authorization"

type ApiClient struct {
	BaseURL  string
	ClientID string
	Token    string
	IsDebug  bool
	Timeout  time.Duration

	AuthData  map[string]any
	UserData  map[string]any
	Axios *http.Client
	Headers    map[string]string
	IsLoginInProgress bool
	EventMap  map[string]*EventSubscription
	mu         sync.Mutex
	Log       utils.Log
}

type EventSubscription struct {
	onMessage   func(eventsource.Event)
	onEvents    map[string]func(eventsource.Event)
	onError     func(error)
	EventSource *eventsource.Stream
}

func NewEventSubscription(stream *eventsource.Stream) *EventSubscription {
	return &EventSubscription{
		onEvents:    make(map[string]func(eventsource.Event)),
		EventSource: stream,
	}
}

type Response struct {
	*http.Response
	Data map[string]any
}

var instance *ApiClient
var once sync.Once

func GetApiClient() *ApiClient {
	return instance
}

func Init(baseURL string, timeout time.Duration, isDebug bool) {
	once.Do(func() {
		// global.EventSource = nodeEventSource;
		// paramsSerializer: params => {
        //   return qs.stringify(params, { arrayFormat: "comma" });
        // },
        // headers: {"X-CLIENT-ID": apiClientUUID },

        if timeout == 0 {
        	timeout = 10000 * time.Millisecond
        }

        clientID := utils.CreateUUID()

		instance = &ApiClient{
			BaseURL:           baseURL,
			ClientID:          clientID,
			Timeout:           timeout,
			Axios:             &http.Client{
				Timeout: timeout,
			},
			Headers:           map[string]string{
				"X-CLIENT-ID": clientID,
			},
			EventMap:          make(map[string]*EventSubscription),
			IsDebug:           isDebug,
			IsLoginInProgress: false,
		}
	})
}

func GetSubscriptions(topic string) *EventSubscription {
	return GetApiClient().EventMap[topic]
}

func (c *ApiClient) Lock() {
	c.mu.Lock()
}

func (c *ApiClient) Unlock() {
	c.mu.Unlock()
}

func (c *ApiClient) SetLocalToken(data map[string]any) {
	var err error

	accessToken := utils.GetStr(data, "accessToken")

	c.Token = accessToken
	c.AuthData = data
	c.UserData, err = utils.DecodeJWT(accessToken)
	if err != nil {
		c.Log.Error("Failed to decode JWT: %s", err)
	}
	c.Headers[AUTH_HEADER] = "Bearer " + accessToken
}

func (c *ApiClient) DeleteLocalToken() {
	c.AuthData = make(map[string]any)
	c.UserData = make(map[string]any)
	delete(c.Headers, AUTH_HEADER)
}

func (c *ApiClient) Post(uri string, bodyJson map[string]any) (*Response, error) {
	return c.Request("POST", uri, bodyJson)
}

func (c *ApiClient) Put(uri string, bodyJson map[string]any) (*Response, error) {
	return c.Request("PUT", uri, bodyJson)
}

func (c *ApiClient) Get(uri string, bodyJson map[string]any) (*Response, error) {
	return c.Request("GET", uri, bodyJson)
}

func (c *ApiClient) Request(method, uri string, bodyJson map[string]any) (*Response, error) {
	uri = c.BaseURL + uri

	resp := &Response{
		Data: make(map[string]any),
	}

	body, err := json.Marshal(bodyJson)
	if err != nil {
		return resp, err
	}

	b := bytes.NewReader(body)
	req, err := http.NewRequest(method, uri, b)
	if err != nil {
		return resp, err
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", "application/json")

	resp.Response, err = c.Axios.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Response.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&resp.Data); err != nil {
		return resp, errors.Wrap(err, "failed to decode token response")
	}

	return resp, err
}

func (es *EventSubscription) AddEventListener(topic string, fn func(e eventsource.Event)) {
	es.onEvents[topic] = fn
}

func (es *EventSubscription) OnMessage(fn func(e eventsource.Event)) {
	es.onMessage = fn
}

func (es *EventSubscription) OnError(fn func(error)) {
	es.onError = fn
}

func (es *EventSubscription) DispatchError(e error) {
	if es.onError != nil {
		es.onError(e)
	}
}

func (es *EventSubscription) DispatchMessage(e eventsource.Event) {
	es.onMessage(e)
}

func (es *EventSubscription) DispatchEvent(id string, e eventsource.Event) {
	if fn, ok := es.onEvents[id]; ok {
		fn(e)
	}
}
