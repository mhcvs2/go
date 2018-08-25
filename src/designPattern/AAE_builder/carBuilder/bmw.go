package carBuilder

import "designPattern/builder/car"

type BMWBuilder struct {
	bmw *car.CarModel
}

func NewBMWBuilder() *BMWBuilder {
	bb := BMWBuilder{
		bmw:&car.CarModel{
			Specific:car.BMWModel{},
		},
	}
	return &bb
}

func (b BMWBuilder)GetCarModel() *car.CarModel {
	return b.bmw
}

func (b BMWBuilder) SetSequence(actions []car.ActionType) {
	b.bmw.SetSequence(actions)
}
