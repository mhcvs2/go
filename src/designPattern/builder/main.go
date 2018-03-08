package main

import "designPattern/builder/director"

func main() {
	d := director.NewDirector()

	a := d.GetABenzModel()
	a.Run()

	b := d.GetBBenzModel()
	b.Run()

	c := d.GetABMWModel()
	c.Run()
}
