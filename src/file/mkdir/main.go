package main

import (
	"os"
	"fmt"
)

func t1() {
	err := os.MkdirAll("/tmp/aa/bb/cc", os.ModePerm)
	if err == nil {
		fmt.Println("mkdir all success")
	} else {
		panic(err)
	}
}

func main() {
	t1()
}
