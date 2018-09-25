package main

import "fmt"

type ITest interface {
	Test()
}

//------------------

type T1 struct {
}

func (t T1) Test() {
	fmt.Println("t1")
}

//--------------------

func d1(t ITest) {
	t.Test()
}

func t1() {
	a := T1{}
	d1(a)
	b := &T1{}
	d1(b)
}

func main() {
	t1()
}