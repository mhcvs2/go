package main

import (
	"stack"
	"fmt"
)

func main() {
	var hstack stack.Stack
	hstack.Push("asd")
	hstack.Push(3.33)
	hstack.Push([] string{"pin","cdcsd"})
	hstack.Push(-11)
	for{
		item, err := hstack.Pop()
		if err != nil{
			break
		}
		fmt.Println(item)
	}
}
