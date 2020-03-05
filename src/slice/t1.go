package main

import (
	"fmt"
	"strings"
)

func t11() {
	s := []int{1,2,3,5}
	fmt.Println(s[0])
	fmt.Println(s[:2])
	s = append(s[:2], s[3:]...)
	fmt.Println(s)
}
//[1 2]
//[1 2 5]

type T struct {
	t int
}

func t12() {
	s := []*T{}
	s = append(s, &T{1})
	s = append(s, &T{2})
	t3 := &T{3}
	s = append(s, t3)
	s = append(s, &T{5})
	for _, ss := range s {
		fmt.Println(ss)
	}
	for i, ss := range s {
		if ss == t3 {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	for _, ss := range s {
		fmt.Println(ss)
	}
}

func t13() {
	a := "sadf dfsdf"
	b := strings.Split(a, " ")
	var c string
	c = b[0]
	fmt.Println(c)
}

func t14() {
	s := []int{0,1}
	b := s[0:2]
	fmt.Println(b)
}

func main() {
	t14()
}
