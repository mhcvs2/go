package main

import "fmt"

func t1() {
	data := []byte{'w','q'}
	src := []byte{'s','a'}
	n := copy(data, src)
	fmt.Println(n)
	fmt.Printf("%q", data)
	//2
	//"sa"
}

func main() {
	t1()
}
