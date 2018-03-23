package main

import "designPattern/visitor/employee2"

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

//name is : haha	job: job1
//name is : lala	per: good