package observable

import (
	"fmt"
	"designPattern/observerPattern/observer"
)

type Observable interface {
	AddObserver(obs observer.Obr)
	DeleteObserver(obs observer.Obr)
	NotifyObservers(context string)
}

type IHanFeiZi interface {
	HaveBreakfast()
	HaveFun()
}

type HanFeiZi struct {
	observerList []observer.Obr
}

func (h *HanFeiZi) AddObserver(obs observer.Obr) {
	h.observerList = append(h.observerList, obs)
}

func (h *HanFeiZi) DeleteObserver(obs observer.Obr) {
	for i, o := range h.observerList {
		if obs == o {
			h.observerList = append(h.observerList[:i], h.observerList[i+1:]...)
			break
		}
	}
}

func (h *HanFeiZi) NotifyObservers(context string) {
	for _, o := range h.observerList {
		o.Update(context)
	}
}

func (h *HanFeiZi) HaveBreakfast() {
	fmt.Println("i bave breakfast")
	h.NotifyObservers("han fei zi have breakfast")
}

func (h *HanFeiZi) HaveFun() {
	fmt.Println("i have fun")
	h.NotifyObservers("han fei zi have fun")
}
