package observer

import (
	"sync"
)

type EventListener interface {
	OnEvent(event interface{})
}

type OwnerShip struct {
	Event interface{}
	Owner EventListener
}

type EventDispatcher struct {
	*EventManager
}

func (e *EventDispatcher) Start() {
	if !e.started{
		e.latch = sync.WaitGroup{}
		e.latch.Add(1)
		e.gracefully = false
		e.flag = false
	}
}

func (e *EventDispatcher) Stop(gracefully bool) {
	if e.started{
		e.gracefully = gracefully
		if gracefully {
			e.latch.Wait()
		}

	}
}
//
//func (e *EventDispatcher) Run() {
//	lastTime := time.Now().Second()
//	var now int
//	var event OwnerShip
//	for{
//		if e.IsAlive(){
//			event = <- e.events
//		}
//
//	}
//}

func (e *EventDispatcher) IsAlive() bool {
	if !e.started {
		return false
	}
	if e.flag && e.wakeup {
		e.flag = false
	}
	return !e.flag
}

type EventManager struct {
	listeners chan EventListener
	events chan OwnerShip
	name string
	eventDispater *EventDispatcher
	started bool
	TriggerNoEvent bool
	Interval int64
	IdleTime int64
	timeout int64
	wakeup bool
	latch sync.WaitGroup
	gracefully bool
	flag bool
}

func NewEventManager(name string, capacity int) *EventManager {
	eventManager := &EventManager{
		name:name,
		events: make(chan OwnerShip, capacity),
		timeout: 1000,
		started: false,
		TriggerNoEvent: false,
		wakeup: false,
	}
	eventManager.eventDispater = &EventDispatcher{
		EventManager: eventManager,
	}
	return eventManager
}



