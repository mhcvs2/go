package main

import "fmt"

type Sc struct {
	F func() bool
}

func test() bool {
	return true
}

func t1() {
	s := Sc{F:test}
	if s.F(){
		fmt.Println("yes")
	}
}

func main() {
	t1()
}
