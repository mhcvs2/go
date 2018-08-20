package main

import (
	"os"
	"io"
)

func t1() {
	dest := os.Stdout
	src, err := os.Open("/root/github/go/src/io/copy/test.txt")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	written, err := io.Copy(dest, src)
	if err != nil {
		println(err.Error())
	} else {
		println(written)
		println("done")
	}
}
//sfasdf
//haha
//hello
//world
//...
//33
//en???done

func main() {
	t1()
}
