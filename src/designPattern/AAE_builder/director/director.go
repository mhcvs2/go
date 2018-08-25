package director

import (
	"designPattern/AAE_builder/car"
	"designPattern/AAE_builder/carBuilder"
)

type Director struct {
	sequence []car.ActionType
	benzBuilder *carBuilder.BenzBuilder
	bmwBuilder *carBuilder.BMWBuilder
}

func NewDirector() *Director {
	return &Director{
		sequence:[]car.ActionType{},
		benzBuilder:carBuilder.NewBenzBuilder(),
		bmwBuilder:carBuilder.NewBMWBuilder(),
	}
}

func (d *Director) GetABenzModel() *car.CarModel {
	actions := []car.ActionType{car.START, car.ALARM}
	d.sequence = actions
	d.benzBuilder.SetSequence(d.sequence)
	return d.benzBuilder.GetCarModel()
}

func (d *Director) GetBBenzModel() *car.CarModel {
	actions := []car.ActionType{car.ALARM, car.STOP}
	d.sequence = actions
	d.benzBuilder.SetSequence(d.sequence)
	return d.benzBuilder.GetCarModel()
}

func (d *Director) GetABMWModel() *car.CarModel {
	actions := []car.ActionType{car.START, car.ALARM}
	d.sequence = actions
	d.bmwBuilder.SetSequence(d.sequence)
	return d.bmwBuilder.GetCarModel()
}
