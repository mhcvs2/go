package main

import (
	"fmt"
	"sync/atomic"
)

func t11() {
	var a, b int32
	a = -288
	b = a & 0x7fffffff
	fmt.Println(b)
}

func t12() {
	var a uint64
	a = 199
	fmt.Println(a)
	atomic.AddUint64(&a, 1)
	defer fmt.Println(a)
	defer atomic.AddUint64(&a, ^uint64(0))
	fmt.Println(a)
}



func main() {
	t12()
}