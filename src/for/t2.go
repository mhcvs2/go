package main

import "fmt"

func t1() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func main() {
	t1()
}
