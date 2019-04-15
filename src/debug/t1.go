package main

import "fmt"

func print(s string) {
	fmt.Println("====")
	fmt.Println(s)
}

func t1() {
	a := "a"
	b := "b"
	c := a + b
	print(c)
}

func main() {
	t1()
}