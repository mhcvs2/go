package main

import (
	"fmt"
	"strconv"
	"math/rand"
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

func t3() {
	a := make(map[string]bool)
	a["a"] = true
	if v, ok := a["k"]; ok{
		fmt.Println("yes")
		fmt.Println(v)
		fmt.Println(ok)
	} else{
		fmt.Println(v)
		fmt.Println(ok)
	}

	fmt.Println(a["d"])

}

func t4() {
	seed := rand.Int63n(4)
	fmt.Println(seed)
}



func main() {
	t4()
}
