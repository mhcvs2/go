package main

import (
	"reflect"
	"fmt"
)

type Test struct {
	A string
}

func t11() {
	t := &Test{A:"haha"}
	value := reflect.ValueOf(t)
	fmt.Println(value.Interface().(*Test).A)
}

func main() {
	t11()
}