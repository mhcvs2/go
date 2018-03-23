package main

import "designPattern/visitor/employee1"

func main() {
	c := employee1.NewCommonEmployee()
	c.SetName("haha")
	c.SetJob("job1")

	m := employee1.NewManager()
	m.SetName("lala")
	m.SetPer("good")

	v := new(employee1.Visitor)

	c.Accept(v)
	m.Accept(v)
}
