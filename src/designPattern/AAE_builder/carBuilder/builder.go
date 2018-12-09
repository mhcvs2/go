package carBuilder

import "designPattern/AAE_builder/car"

type CarBuilder interface {
	SetSequence(actions []car.ActionType)
	GetCarModel() *car.ICarCommon
}

