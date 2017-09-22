package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func hello() {
	fmt.Println("hello world")
}

func prints(i int) string {
	fmt.Println("i =", i)
	return strconv.Itoa(i)
}

func t1() {
	funmmap := make(map[string]func())
	funmmap["hello"] = hello
	funmmap["hello"]()
}

func t2() {
	hl := hello
	fv := reflect.ValueOf(hl)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)
}

func t3() {
	fv := reflect.ValueOf(prints)
	params := make([]reflect.Value, 1)                 // 参数
	params[0] = reflect.ValueOf(20)                    // 参数设置为20
	rs := fv.Call(params)                              // rs作为结果接受函数的返回值
	fmt.Println("result:", rs[0].Interface().(string)) // 当然也可以直接是 rs[0].Interface()
}

func main()  {
	t3()
}
