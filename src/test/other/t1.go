package main

import (
	"fmt"
	"path/filepath"
	"os"
	"bufio"
)

func t11() {
	fmt.Println(filepath.Join("aaa","bb"))
}

func t12() {
	f, _ := os.Open("/root/github/go/src/test/other/test.txt")
	r := bufio.NewReader(f)
	fmt.Println(r.ReadLine())
}

func t13() {
	i := 0
	fmt.Println(i)
}


func main() {
	t12()
}
