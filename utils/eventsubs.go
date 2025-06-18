package utils

import (
	"github.com/donovanhide/eventsource"
	"github.com/pkg/errors"
)

type EventStream = eventsource.Stream
type Event = eventsource.Event

type EventSubscription struct {
	onMessage   func(Event)
	onEvents    map[string]func(Event)
	onError     func(error)
	EventSource *EventStream
	isRegist    bool
}

func NewEventSubscription(url, lastEventId string) (*EventSubscription, error) {
	stream, err := eventsource.Subscribe(url, lastEventId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to subscribe")
	}

	return &EventSubscription{
		onEvents:    make(map[string]func(Event)),
		EventSource: stream,
	}, nil
}

func (es EventSubscription) IsRegist() bool {
	return es.isRegist
}

func (es *EventSubscription) SetRegist(state bool) {
	es.isRegist = state
}

func (es *EventSubscription) AddEventListener(topic string, fn func(e Event)) {
	es.onEvents[topic] = fn
}

func (es *EventSubscription) OnMessage(fn func(e Event)) {
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

func (es *EventSubscription) DispatchMessage(e Event) {
	es.onMessage(e)
}

func (es *EventSubscription) DispatchEvent(id string, e Event) {
	if fn, ok := es.onEvents[id]; ok {
		fn(e)
	}
}

func (es *EventSubscription) EventLoop() {
	// fmt.Println("EventLoop start...")
	for {
		select {
		case e := <-es.EventSource.Events:
			if e == nil {
				es.DispatchEvent("error", nil)
				//fmt.Println("EventLoop end...")
				return
			}

			switch e.Event() {
			case "error":
				es.DispatchError(errors.New(e.Data()))
			case "message", "":
				es.DispatchMessage(e)
			default:
				// fmt.Printf("EventLoop event... [%s] %+v\n", e.Event(), e)
				es.DispatchEvent(e.Event(), e)
			}
		case err := <-es.EventSource.Errors:
			if err != nil {
				es.DispatchError(err)
				// fmt.Println("EventLoop end...")
				return
			}
		}
	}
}
