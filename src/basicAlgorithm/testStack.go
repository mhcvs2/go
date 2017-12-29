package main

import (
	"basicAlgorithm/structure"
	"fmt"
)

func ts1() {
	stack := structure.NewStack(10)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	fmt.Println(stack.Length())
	fmt.Println(stack.Top())
	fmt.Println(stack.Length())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Length())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Length())
}

func ts2() {
	stack := structure.NewQueue(10)
	stack.Push(5)
	stack.Push(4)
	stack.Push(3)
	stack.Push(2)
	fmt.Println(stack.Length())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Length())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Length())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Length())
}

func main() {
	ts1()
}
