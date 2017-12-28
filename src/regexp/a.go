package main

import (
	"regexp"
	"fmt"
)


var text = "The quick brown fox jumped..."

func a1() {
	obj := regexp.MustCompile("fox")
	fmt.Println(obj.FindString(text))
}

func main() {
	a1()
}
