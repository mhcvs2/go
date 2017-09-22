package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func t1() {
	a, _ := os.Getwd()
	fmt.Println(a)
}

func t2() {
	a := filepath.Dir("/aa/bb/cc")
	fmt.Println(a)
}

func main() {
	t2()
}
