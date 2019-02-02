package main

import (
	"designPattern/AAP_observer/c_event/observer"
	"fmt"
	"time"
)

type KillTaskListener struct {}

func (l *KillTaskListener) OnEvent(event interface{}) {
	fmt.Println("KillTaskListener")
}

func (l *KillTaskListener) Handles(event interface{}) bool {
	return true
}

//===========================

type DispatchListener struct {}

func (l *DispatchListener) OnEvent(event interface{}) {
	fmt.Println("KillTaskListener")
}

func (l *DispatchListener) Handles(event interface{}) bool {
	return true
}

//========================

type DispatchEventManager struct {
	*observer.EventManager
	dispatcher *Dispatcher
}

func NewDispatchEventManager(dispatcher *Dispatcher) *DispatchEventManager {
	res := &DispatchEventManager{observer.NewEventManager("DispatchEventManager", 50),
		dispatcher}
	res.EventManager.Specific = res
	return res
}

func (e *DispatchEventManager) Publish(event interface{}) {
	e.EventManager.Publish(event)
}

func (e *DispatchEventManager) OnIdle() {
	e.dispatcher.Dispatch()
}

//============
type Dispatcher struct {
	dispatchEventmanager *DispatchEventManager
	killtaskEventManager *observer.EventManager
}

func NewDispatcher() *Dispatcher {
	dispatcher := &Dispatcher{}
	dispatchEventmanager := NewDispatchEventManager(dispatcher)
	killtaskEventManager := observer.NewEventManager("killtaskEventManager", 50)

	dispatchListener := &DispatchListener{}
	dispatchEventmanager.AddListener(dispatchListener)
	dispatchEventmanager.IdleTime = 3
	dispatchEventmanager.Interval = 5

	killTaskListener := &KillTaskListener{}
	killtaskEventManager.AddListener(killTaskListener)
	killtaskEventManager.Interval = 3
	killtaskEventManager.TriggerNoEvent = true

	dispatcher.dispatchEventmanager = dispatchEventmanager
	dispatcher.killtaskEventManager = killtaskEventManager

	//dispatchEventmanager.Start()
	killtaskEventManager.Start()
	return dispatcher
}

func (d *Dispatcher) Dispatch() {
	fmt.Println("Dispatch")
}

func main() {
	d := NewDispatcher()
	fmt.Println(d)
	time.Sleep(time.Second * 50)
}