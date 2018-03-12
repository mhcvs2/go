package main

import (
	"fmt"
)

type Cloneable interface {
	Clone() interface{}
}

type A struct {
	a int
	b []int
}

func (c *A) Clone() *A {
	new_obj := *c
	return &new_obj
}

func t1() {
	var a *A = &A{}
	a.a = 100
	b := a.Clone()
	fmt.Println("------------")
	fmt.Println(a.a)
	fmt.Println(b.a)
	b.a = 999
	fmt.Println("------------")
	fmt.Println(a.a)
	fmt.Println(b.a)


}
//------------
//100
//100
//------------
//100
//999

func t2() {
	var a *A = &A{}
	a.a = 100
	a.b = []int{1,2,3}
	b := a.Clone()
	fmt.Println("------------")
	fmt.Println(a.b)
	fmt.Println(b.b)
	b.b = append(b.b, 100)
	fmt.Println("------------")
	fmt.Println(a.b)
	fmt.Println(b.b)
	fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", b)
}
//------------
//[1 2 3]
//[1 2 3]
//------------
//[1 2 3]
//[1 2 3 100]
//&{64 [1 2 3]}
//&{64 [1 2 3 64]}

func t3() {
	var a *A= &A{}   /* 声明实际变量 */
	var ip **A        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", **ip )
}
//a 变量的地址是: c42000c028
//ip 变量的存储地址: c42000c028
//*ip 变量的值: {0 []}

func main() {
	t2()
}