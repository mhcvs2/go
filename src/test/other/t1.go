package main

import (
	"fmt"
	"path/filepath"
)

func t11() {
	fmt.Println(filepath.Join("aaa","bb"))
}

func t12() {
	fmt.Println(10 ^ 1)
}


func main() {
	t12()
}
