package main

import (
	"fmt"
)

func test(cur interface{}) {
	if cur == nil {
		fmt.Println("cur is nil")
	} else {
		fmt.Println("cur is not nil")
	}
}

func main() {
	var cur interface{}
	test(cur)
}
