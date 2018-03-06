package main

import "fmt"

/*
不能一个interface 跨两个struct
*/

type Type1 interface {
	A(a string) string
	B(b int) int
}

type Base struct {
	m string
}

func (base *Base)A(a string) string {
	fmt.Println("A in base")
	return a
}

func (base *Base)A2(a string) string {
	fmt.Println("A2 in base")
	return a
}

func (base *Base)A3() string {
	fmt.Println("A3 in base")
	return base.m
}

type Concront struct {
	Base
}

func (con *Concront)A(a string) string {
	fmt.Println("A in concront")
	return con.Base.A(a)
}

func (con *Concront)B(b int) int {
	return b
}

func test(tp Type1) {
	fmt.Println(tp.A("haha"))
	fmt.Println(tp.B(2222))
}

func main() {
	c := &Concront{Base{m:"ddd"}}
	test(c)
	c.A2("lala")
	fmt.Println(c.A3())
}