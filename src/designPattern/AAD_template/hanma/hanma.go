package hanma

import "fmt"

type IHanma interface {
	Start()
	Stop()
	Alarm()
	Run()
}

type HanmaModel struct {
	Specific IHanma
}

func (h HanmaModel) Run() {
	if h.Specific != nil {
		h.Specific.Start()
		h.Specific.Alarm()
		h.Specific.Stop()
	} else {
		fmt.Println("error")
	}
}