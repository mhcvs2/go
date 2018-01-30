package main

import "fmt"

func t11() {
	var a, b int32
	a = -288
	b = a & 0x7fffffff
	fmt.Println(b)
}



func main() {
	t11()
}