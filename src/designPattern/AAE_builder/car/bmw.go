package car

import "fmt"

type BMWModel struct {
	CarModel
}

func (h BMWModel) Start() {
	fmt.Println("BMW start")
}

func (h BMWModel) Stop() {
	fmt.Println("BMW Stop")
}

func (h BMWModel) Alarm() {
	fmt.Println("BMW Alarm")
}
