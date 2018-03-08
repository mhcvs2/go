package hanma

import "fmt"

type HanmaH2Model struct {
	HanmaModel
}

func (h HanmaH2Model) Start() {
	fmt.Println("h2 start")
}

func (h HanmaH2Model) Stop() {
	fmt.Println("h2 Stop")
}

func (h HanmaH2Model) Alarm() {
	fmt.Println("h2 Alarm")
}