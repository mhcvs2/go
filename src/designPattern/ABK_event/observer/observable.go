package observer

import (
	"fmt"
)

type IHanFeiZi interface {
	HaveBreakfast()
	HaveFun()
}

type HanFeiZi struct {
	*Observable
}

func NewHanFeiZi() *HanFeiZi {
	return &HanFeiZi{
		NewObservable(),
	}
}

func (h *HanFeiZi) HaveBreakfast() {
	fmt.Println("i bave breakfast")
	h.SetChanged()
	h.NotifyObservers("han fei zi have breakfast")
}

func (h *HanFeiZi) HaveFun() {
	fmt.Println("i have fun")
	h.SetChanged()
	h.NotifyObservers("han fei zi have fun")
}
