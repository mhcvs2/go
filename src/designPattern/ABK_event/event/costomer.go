package event

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
}

type EventCustomer struct {
	customTypes []EventCustomType
}

func (s *EventCustomer)AddCustomType(t EventCustomType) {
    s.customTypes = append(s.customTypes, t)
}

func (s *EventCustomer)CustomTypes() []EventCustomType {
    return s.customTypes
}

