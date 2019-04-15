package main

import "fmt"

func add(a, b int) int  {
	return a + b
}

func main() {
	a := []int{1,2,3}
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println(add(1,456))
}