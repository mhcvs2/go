package main

import "fmt"

type Test interface{
	A()
}

type Mhc struct {

}

func (m Mhc) A() {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
}

//---------------------------------------------

type T2 struct {
	Test
}

func (t T2) A() {
	t.Test.A()
}

type T3 struct {
	Test
}

func (t T3) B() {
	t.Test.A()
}

func main() {
	m := &Mhc{}
	t2 := &T2{m}
	t2.A()

	t3 := &T3{m}
	t3.A()
	t3.B()
}
//aaaaaaaaaaaaaaaaaaaaaa
//aaaaaaaaaaaaaaaaaaaaaa
//aaaaaaaaaaaaaaaaaaaaaa