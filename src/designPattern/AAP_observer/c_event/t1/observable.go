package t1

import (
	"designPattern/AAP_observer/c_event/observer"
	"fmt"
)

type HanFeiZiEventType int

const (
	HAVEFUN HanFeiZiEventType = iota
	HAVEBREAKFAST
)

type Lisi struct {}

func (l *Lisi) OnEvent(event interface{}) {
	fmt.Println("lisi: watch hanfeizi act")
	switch event {
	case HAVEFUN:
		fmt.Println("hanfeizi have fun")
	case HAVEBREAKFAST:
		fmt.Println("hanfeizi have breakfast")
	}
}

func (l *Lisi) Handles(event interface{}) bool {
	return  event == HAVEFUN || event == HAVEBREAKFAST
}

//================================================================
type IHanFeiZi interface {
	HaveBreakfast()
	HaveFun()
}

type HanFeiZi struct {
	eventManager *observer.EventManager
	eventListener observer.EventListener
}

func NewHanFeiZi() *HanFeiZi {
	han := &HanFeiZi{
		eventManager: observer.NewEventManager("hanfeizi", 50),
		eventListener: &Lisi{},
	}
	han.eventManager.AddListener(han.eventListener)
	han.eventManager.Start()
	return han
}

func (h *HanFeiZi) HaveBreakfast() {
	fmt.Println("i bave breakfast")
	if ok := h.eventManager.Add(HAVEBREAKFAST, h.eventListener); !ok{
		fmt.Println("report failed")
	}
}

func (h *HanFeiZi) HaveFun() {
	fmt.Println("i have fun")
	if ok := h.eventManager.Add(HAVEFUN, h.eventListener); !ok{
		fmt.Println("report failed")
	}
}

func (h *HanFeiZi) GetEventManager() *observer.EventManager {
	return h.eventManager
}