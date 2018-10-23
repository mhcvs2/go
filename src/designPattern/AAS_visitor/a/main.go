package main

import (
	"designPattern/AAS_visitor/a/employee"
)

func main() {
	c := employee.NewCommonEmployee()
	c.SetName("haha")
	c.SetJob("job1")

	m := employee.NewManager()
	m.SetName("lala")
	m.SetPer("good")

	v := new(employee.Visitor)

	c.Accept(v)
	m.Accept(v)
}
