package hanma

import "fmt"

type HanmaH1Model struct {
	HanmaModel
}

func (h HanmaH1Model) Start() {
	fmt.Println("h1 start")
}

func (h HanmaH1Model) Stop() {
	fmt.Println("h1 Stop")
}

func (h HanmaH1Model) Alarm() {
	fmt.Println("h1 Alarm")
}