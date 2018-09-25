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

type T2 struct {
	Son ITest
}

func (t T2) Test2() {
	t.Son.Test()
}

func t1() {
	a := T1{}
	aa := T2{a}
	aa.Test2()
	b := &T1{}
	bb := T2{b}
	bb.Test2()
}

func main() {
	t1()
}

//t1
//t1