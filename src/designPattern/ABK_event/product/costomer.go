package product

import "fmt"

type EventCustomType int

const (
	NEW EventCustomType = iota
	DEL
	EDIT
	CLONE
)

type ICustomer interface {
	AddCustomType(t EventCustomType)
	CustomTypes() []EventCustomType
	Exec(event ProductEvent)
}

type EventCustomer struct {
	customTypes []EventCustomType
	spec ICustomer
}

func NewEventCustomer(eventCustomType EventCustomType, spec ICustomer) *EventCustomer {
	res := &EventCustomer{customTypes: []EventCustomType{}, spec: spec}
	res.AddCustomType(eventCustomType)
	return res
}

func (s *EventCustomer) Exec(event ProductEvent) {
	panic("implement me")
}

func (s *EventCustomer)AddCustomType(t EventCustomType) {
    s.customTypes = append(s.customTypes, t)
}

func (s *EventCustomer)CustomTypes() []EventCustomType {
    return s.customTypes
}

type Beggar struct {
	*EventCustomer
}

func NewBeggar() *Beggar {
	res := &Beggar{}
	res.EventCustomer = NewEventCustomer(DEL, res)
	return res
}

func (s *Beggar) Exec(event ProductEvent) {
	p := event.Source()
	eventType := event.EventType()
	fmt.Printf("乞丐处理事件: %s 销毁, 事件类型: %d\n", p.Name(), eventType)
}

type Commoner struct {
	*EventCustomer
}

func NewCommoner() *Commoner {
	res := &Commoner{}
	res.EventCustomer = NewEventCustomer(NEW, res)
	return res
}

func (s *Commoner) Exec(event ProductEvent) {
	p := event.Source()
	eventType := event.EventType()
	fmt.Printf("平民处理事件: %s 新建, 事件类型: %d\n", p.Name(), eventType)
}

type Nobleman struct {
	*EventCustomer
}

func NewNobleman() *Nobleman {
	res := &Nobleman{}
	res.EventCustomer = NewEventCustomer(CLONE, res)
	res.AddCustomType(EDIT)
	return res
}

func (s *Nobleman) Exec(event ProductEvent) {
	p := event.Source()
	eventType := event.EventType()
	if int(eventType) == int(CLONE) {
		fmt.Printf("贵族处理事件: %s 克隆, 事件类型: %d\n", p.Name(), eventType)
	} else {
		fmt.Printf("贵族处理事件: %s 修改, 事件类型: %d\n", p.Name(), eventType)
	}

}