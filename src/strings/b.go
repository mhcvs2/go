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

func b3() {
	containerID := "docker://d3168008a2bf371703c970bf0ca48208b6d0f8db3d3da9ee77c50ad11ae8aabb"
	res := strings.Split(containerID,"docker://")[1]
	fmt.Println(res)
}

func main() {
	b3()
}