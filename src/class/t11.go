package main

import "fmt"

type interface1 interface {
	Id() int
}

type interface2 interface {
	Id() int
	Name() string
}

type testStruct struct {
	id int
	name string
}

func (t testStruct) Id() int {
	return t.id
}

func (t testStruct) Name() string {
	return t.name
}

type type1 func() interface2

func testFunc() interface2 {
	s := testStruct{
		id:1,
		name:"mhc",
	}
	return &s
}

func t11() {
	var testF type1 = testFunc
	testF2 := func() interface1 {
		return testF()
	}
	fmt.Println(testF2().Id())
}

func main() {
	t11()
}
