package observer

import (
	"sync/atomic"
	"time"
)

type IQueue interface {
	Get(timeout int64) interface{}
	Put(item interface{}, timeout int64) bool
	GetBlock() interface{}
	IsEmpty() bool
	Size() int
}

type EventQueue struct {
	capacity int
	data chan interface{}
	size int32
}

func (e *EventQueue) IsEmpty() bool {
	return e.size == 0
}

func (e *EventQueue) Size() int {
	return int(e.size)
}

func (e *EventQueue) GetBlock() interface{} {
	if ret, ok := <- e.data; ok {
		atomic.AddInt32(&e.size, -1)
		return ret
	}
	return nil
}

func (e *EventQueue) Get(timeout int64) interface{} {

	select {
	case ret := <- e.data:
		atomic.AddInt32(&e.size, -1)
		return ret
	case <-time.After(time.Second * time.Duration(timeout)):
		return nil
	}
}

func (e *EventQueue) Put(item interface{}, timeout int64) bool {
	select {
	case e.data <- item:
		atomic.AddInt32(&e.size, 1)
		return true
	case <- time.After(time.Second * time.Duration(timeout)):
		return false

	}
}



func NewEventQueue(capacity int) *EventQueue {
	return &EventQueue{
		capacity: capacity,
		data: make(chan interface{}, capacity),
		size: 0,
	}
}

