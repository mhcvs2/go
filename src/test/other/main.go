package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func t1() {
	data := []byte{'w','q'}
	src := []byte{'s','a'}
	n := copy(data, src)
	fmt.Println(n)
	fmt.Printf("%q", data)
	//2
	//"sa"
}

func lenOfString(str string) string  {
	l := len([]rune(str))
	return strconv.Itoa(l)
}

func t2() {
	fmt.Println(lenOfString("hello"))
	reflect.ValueOf()
}



func main() {
	t2()
}
