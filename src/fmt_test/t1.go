package main

import (
	"fmt"
)

func t11() {
	var a int
	var b int
	fmt.Sscanf("min=2,max=1", "min=%d,max=%d", &a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

//2
//1

func t12() {
	var (
		old = "1.00000023e+06"
		new float64
	)
	n, err := fmt.Sscanf(old, "%f", &new)
	if err != nil {
		fmt.Println(err.Error())
	} else if 1 != n {
		fmt.Println("n is not one")
	}
	fmt.Println(uint64(new))
}
//1000000

func main() {
	t12()
}