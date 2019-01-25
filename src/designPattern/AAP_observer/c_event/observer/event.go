package observer

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type EventListener interface {
	OnEvent(event interface{})
	Handles(event interface{}) bool
}

type OwnerShip struct {
	Event interface{}
	Owner EventListener
}

type IEventManager interface {
	Publish(event interface{})
	OnIdle()
}

//=========================================

type EventManager struct {
	listeners []EventListener
	events IQueue
	name string
	started bool
	TriggerNoEvent bool
	Interval int64
	IdleTime int64
	timeout int64
	wakeup bool
	latch sync.WaitGroup
	gracefully bool
	flag bool
	Specific IEventManager
}

func NewEventManager(name string, capacity int) *EventManager {
	eventManager := &EventManager{
		name:name,
		events: NewEventQueue(capacity),
		timeout: 2,
		started: false,
		TriggerNoEvent: false,
		wakeup: false,
		listeners: make([]EventListener, 0),
	}
	eventManager.IdleTime = 0
	eventManager.Interval = 0
	eventManager.Specific = eventManager
	return eventManager
}

func (e *EventManager) Start() {
	if !e.started{
		e.latch = sync.WaitGroup{}
		e.latch.Add(1)
		e.gracefully = false
		e.flag = false
		e.started = true
	}
	go e.Run()
}

func (e *EventManager) Stop(gracefully bool) {
	if e.started{
		e.started = false
		e.gracefully = gracefully
		if gracefully {
			e.latch.Wait()
		}

	}
}

func (e *EventManager) getEvent(block bool) *OwnerShip {
	var data interface{}
	if block {
		data = e.events.GetBlock()
	} else {
		data = e.events.Get(e.timeout)
	}
	if data != nil {
		return data.(*OwnerShip)
	} else {
		return nil
	}
}

func (e *EventManager) Run() {
	fmt.Println("run")
	fmt.Println(e.Interval)
	fmt.Println(e.IdleTime)
	lastTime := time.Now().Second()
	var now int
	var event *OwnerShip
	for{
		if e.IsAlive(){
			event = e.getEvent(false)
		}
		if !e.IsAlive(){
			if !e.gracefully {
				break
			}
			if event == nil {
				if !e.flag {
					event = e.getEvent(false)
				} else {
					e.getEvent(true)
				}
			}
			if event == nil {
				break
			}
		}
		if event != nil {
			if e.IdleTime > 0 {
				lastTime = time.Now().Second()
			}
			if e.Interval > 0 {
				currents := make([]*OwnerShip, 0)
				for ; event != nil; {
					currents = append(currents, event)
					e.getEvent(false)
				}
				e.Specific.Publish(currents)
			} else {
				e.Specific.Publish(event)
			}
		} else {
			if e.TriggerNoEvent {
				e.Specific.Publish(nil)
			}
			if e.IdleTime > 0 {
				now = time.Now().Second()
				fmt.Println(now - lastTime)
				if int64(now - lastTime) > e.IdleTime {
					lastTime = now
					e.Specific.OnIdle()
					if !e.events.IsEmpty() {
						continue
					}
				}
			}
		}
		if e.Interval > 0 && e.IsAlive() {
			time.Sleep(time.Second * time.Duration(e.Interval))
		}

	}
	e.latch.Done()
}

func (e *EventManager) IsAlive() bool {
	if !e.started {
		return false
	}
	if e.flag && e.wakeup {
		e.flag = false
	}
	return !e.flag
}

func (e *EventManager) Publish(event interface{}) {
	var eventData *OwnerShip
	if reflect.ValueOf(event).Kind() == reflect.Slice {

		eventSlice := event.([]*OwnerShip)
		if event == nil || len(eventSlice) == 0 {
			return
		}
		eventData = eventSlice[0]
	} else if event != nil {
		eventData = event.(*OwnerShip)
	} else {
		eventData = nil
	}
	if eventData != nil && eventData.Owner != nil {
		eventData.Owner.OnEvent(eventData.Event)
	} else {
		var data interface{}
		if eventData == nil {
			data = nil
		} else {
			data = eventData.Event
		}
		for _, listener := range e.listeners {
			if listener.Handles(data) {
				listener.OnEvent(data)
			}
		}
	}
}

func (e *EventManager) OnIdle() {}

func (e *EventManager) AddListener(listener EventListener) bool {
	e.listeners = append(e.listeners, listener)
	return true
}

func (e *EventManager) Add(event interface{}, owner EventListener) bool {
	if event == nil {
		return false
	}
	return e.events.Put(&OwnerShip{event, owner}, 1)
}
