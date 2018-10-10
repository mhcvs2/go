package main

import (
	"fmt"
	"reflect"
)

func t1()  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}
//defer panic
//recover的是最后一个panic

func t2()  {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("++++")
			f:=err.(func()string)
			fmt.Println(err,reflect.TypeOf(err).Kind().String())
			fmt.Println(f())
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return  "defer panic"
		})
	}()
	panic("panic")
}
//++++
//0x488000 func
//defer panic

func main() {
	t2()
}
