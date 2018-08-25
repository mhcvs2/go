package car

import "fmt"

type ActionType int

const (
	START ActionType = iota
	STOP
	ALARM
)

type ICarCommon interface {
	Run()
	SetSequence(actions []ActionType)
}

type ICar interface {
	Start()
	Stop()
	Alarm()
}

type CarModel struct {
	actions []ActionType
	Specific ICar
}

func (h CarModel) Run() {
	if h.Specific != nil {
		for _, action := range h.actions {
			switch action {
			case START:
				h.Specific.Start()
			case STOP:
				h.Specific.Stop()
			case ALARM:
				h.Specific.Alarm()
			}
		}
	} else {
		fmt.Println("error")
	}
}

func (h *CarModel) SetSequence(actions []ActionType) {
	h.actions = actions
}