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

func t2() {
	m := make(map[int]map[int]string)
	if _,ok := m[2]; !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "good"
	a, ok := m[2][1]
	fmt.Println(a, ok)
	c, ok := m[2]
	fmt.Println(c, ok)

}

func main() {
	t2()
}