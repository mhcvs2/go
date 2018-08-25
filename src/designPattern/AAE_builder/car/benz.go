package car

import "fmt"

type BenzModel struct {
	CarModel
}

func (h BenzModel) Start() {
	fmt.Println("Benz start")
}

func (h BenzModel) Stop() {
	fmt.Println("Benz Stop")
}

func (h BenzModel) Alarm() {
	fmt.Println("Benz Alarm")
}
