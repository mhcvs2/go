package product

import (
	"designPattern/ABK_event/observer"
	"fmt"
	"sync"
)

type ProductEventType int

const (
	NEW_PRODUCT ProductEventType = iota
	DEL_PRODUCT
	EDIT_PRODUCT
	CLONE_PRODUCT
)

var dispatch *EventDispatch
var once sync.Once

func GetEventDispatch() *EventDispatch {
	once.Do(func() {
		dispatch = NewEventDispatch()
	})
	return dispatch
}

type ProductEvent struct {
	*observer.Observable
	source Product
	eventType ProductEventType
}

func NewProductEvent(source Product, eventType ProductEventType) *ProductEvent {
	res := &ProductEvent{
		source: source,
		eventType: eventType,
	}
	res.Observable = observer.NewObservable(res)
	res.NotifyEventDispatch()
	return res
}

func (s *ProductEvent)Source() Product {
    return s.source
}

func (s *ProductEvent)EventType() ProductEventType {
    return s.eventType
}

func (s *ProductEvent)NotifyEventDispatch() {
	s.AddObserver(GetEventDispatch())
    s.SetChanged()
    s.NotifyObservers(s.source)
}

type EventDispatch struct {
	customers []ICustomer
}

func (s *EventDispatch) Update(o observer.IObservable, arg interface{}) {
	p := arg.(Product)
	fmt.Println(p.Name())
	event := o.(*ProductEvent)
	for _, customer := range s.customers {
		for _, eventCustomerType := range customer.CustomTypes() {
			if int(eventCustomerType) == int(event.EventType()) {
				customer.Exec(*event)
			}
		}
	}

}

func NewEventDispatch() *EventDispatch {
	return &EventDispatch{customers: []ICustomer{}}
}

func (s *EventDispatch)RegisterCustomer(customer ICustomer) {
    s.customers = append(s.customers, customer)
}
