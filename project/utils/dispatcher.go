package utils

import "project/observers"

type EventDispatcher struct {
	observers map[string][]observers.Observer
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		observers: make(map[string][]observers.Observer),
	}
}

func (ed *EventDispatcher) Register(event string, observer observers.Observer) {
	ed.observers[event] =append(ed.observers[event], observer)

}

func (ed *EventDispatcher) Dispatch(event string, data interface{}) {
	observers := ed.observers[event]
	for _, observer := range observers {
		observer.Update(event,data)
	}
}