package main

import (
	"path/filepath"
	"fmt"
)

func t1() {
	path := "/etc/sss/aaa.txt"
	fmt.Println(path)
	fmt.Println(filepath.Ext(path))
	///etc/sss/aaa.txt
	//.txt
}

func main() {
	t1()
}
