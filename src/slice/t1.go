package main

import "fmt"

func t11() {
	s := []int{1,2,3,5}
	s = append(s[:2], s[3:]...)
	fmt.Println(s)
}

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

func main() {
	t12()
}
