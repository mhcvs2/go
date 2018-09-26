package main

import "fmt"

func t1() {
	a := make(map[string]bool)
	a["a"] = false
	a["b"] = false
	fmt.Println(a)
	delete(a, "a")
	fmt.Println(a)
	delete(a, "c")
	fmt.Println(a)
}
//map[a:false b:false]
//map[b:false]
//map[b:false]

func main() {
	t1()
}