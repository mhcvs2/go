package main

import (
	"sort"
	"fmt"
)

type myType []int

func (m myType) Len() int {
	return len(m)
}

func (m myType) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m myType) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func t11() {
	a := myType{5, 3, 2, 8, 1}
	sort.Sort(a)
	fmt.Println(a)
	a2 := sort.Reverse(a)
	fmt.Println(a2)
	sort.Sort(a2)
	fmt.Println(a2)
}
//[1 2 3 5 8]
//&{[1 2 3 5 8]}
//&{[8 5 3 2 1]}

func t12() {
	a := myType{5, 3, 2, 8, 1}
	a2 := sort.Reverse(a)
	fmt.Println(a2)
	sort.Sort(a2)
	fmt.Println(a2)
}
//&{[5 3 2 8 1]}
//&{[8 5 3 2 1]}

func main() {
	t12()
}
