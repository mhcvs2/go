package main

import "fmt"

func defer_call() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	panic("触发异常")
}

func main() {
	defer_call()
}
