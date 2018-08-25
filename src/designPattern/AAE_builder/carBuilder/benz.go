package carBuilder

import "designPattern/AAE_builder/car"

type BenzBuilder struct {
	benz *car.CarModel
}

func NewBenzBuilder() *BenzBuilder {
	bb := BenzBuilder{
		benz:&car.CarModel{
			Specific:car.BenzModel{},
		},
	}
	return &bb
}

func (b BenzBuilder)GetCarModel() *car.CarModel {
	return b.benz
}

func (b BenzBuilder) SetSequence(actions []car.ActionType) {
	b.benz.SetSequence(actions)
}
