package main

import "fmt"

type I1 interface {
	A()
	B() string
}

type T1 struct {
	I1
}

type T2 struct {}

func (t T2)A() {
	fmt.Println("aaa")
}

func (t T2)B() string {
	fmt.Println("bbb")
	return "bbb"
}

//------------------------------------
func main() {
	t := T1{I1:T2{}}
	t.I1.B()
}

