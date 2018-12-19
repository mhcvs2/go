package observable

import (
	"fmt"
	"designPattern/AAP_observer/b/observer"
)

type IHanFeiZi interface {
	HaveBreakfast()
	HaveFun()
}

type HanFeiZi struct {
	*observer.Observable
}

func NewHanFeiZi() *HanFeiZi {
	return &HanFeiZi{
		observer.NewObservable(),
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
