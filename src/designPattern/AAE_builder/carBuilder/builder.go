package carBuilder

import "designPattern/builder/car"

type CarBuilder interface {
	SetSequence(actions []car.ActionType)
	GetCarModel() *car.ICarCommon
}

