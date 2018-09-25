package main

import "fmt"

type ITest interface {
	Test()
	Set(s string)
}

//------------------

type T1 struct {
	Data string
}

func (t T1) Test() {
	fmt.Println(t.Data)
}

func (t *T1) Set(s string) {
	t.Data = s
}

//--------------------

type T2 struct {
	Son ITest
}

func (t T2) Test2() {
	t.Son.Test()
}

func (t *T2) update(s string) {
	t.Son.Set(s)
}

func t1() {
	//a := T1{Data: "a"}
	//aa := T2{a}   //出错
	//aa.Test2()
	b := &T1{}
	bb := T2{b}
	bb.Test2()
}

func main() {
	t1()
}

// 如果实现interface的方法中是 指针类型， 那么给接口变量赋值，只能赋指针类型