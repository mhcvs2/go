package main

import (
	"reflect"
	"fmt"
	"strings"
)

type Test struct {
	A string
}

func t11() {
	t := &Test{A:"haha"}
	value := reflect.ValueOf(t)
	fmt.Println(value.Interface().(*Test).A)
	ts := value.Type().String()
	tss := strings.Split(ts, ".")
	fmt.Println(tss[len(tss)-1])
}

func main() {
	t11()
}