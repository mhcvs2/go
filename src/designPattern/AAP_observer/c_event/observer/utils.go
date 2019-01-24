package observer

import "time"

type IQueue interface {
	Get(timeout int) interface{}
	Put(item interface{}, timeout int) bool
	GetBlock() interface{}
}

type EventQueue struct {
	capacity int
	data chan interface{}
}

func (e *EventQueue) GetBlock() interface{} {
	if ret, ok := <- e.data; ok {
		return ret
	}
}

func (e *EventQueue) Get(timeout int) interface{} {
	select {
	case ret := <- e.data:
		return ret
	case <-time.After(time.Second * time.Duration(timeout)):
		return nil
	}
}

func (e *EventQueue) Put(item interface{}, timeout int) bool {
	select {
	case e.data <- item:
		return true
	case <- time.After(time.Second * time.Duration(timeout)):
		return false

	}
}



func NewEventQueue(capacity int) *EventQueue {
	return &EventQueue{
		capacity: capacity,
		data: make(chan interface{}, capacity),
	}
}

