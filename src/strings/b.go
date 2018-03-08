package main

import (
	"fmt"
	"strings"
	"bytes"
)

func b1() {
	a := []string{"1", "2", "e"}
	fmt.Println(strings.Join(a, ""))
	fmt.Println(strings.Join(a, "-"))
}
//12e
//1-2-e

func b2() {
	var buffer bytes.Buffer
	buffer.WriteString("hahah")
	buffer.WriteString("llll")
	fmt.Println(buffer.String())
}

func main() {
	b2()
}