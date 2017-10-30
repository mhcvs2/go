package main

import (
	"fmt"
	"strconv"
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

type S struct {
	A string
	B bool
}

func t2() {
	s := S{}
	if s.B {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}



func main() {
	t2()
}
