package main

import "designPattern/AAE_builder/director"

func main() {
	d := director.NewDirector()

	a := d.GetABenzModel()
	b := d.GetBBenzModel()
	c := d.GetABMWModel()

	a.Run()
	b.Run()
	c.Run()
}
