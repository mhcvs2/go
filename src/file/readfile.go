package main

import (
	"io/ioutil"
	"fmt"
)

func Readfile() {
	path := "/root/github/go/src/file/a.txt"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {
	Readfile()
}