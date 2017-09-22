package main

import "fmt"

func getMaxLen(strs []string) int {
	max := 0
	for _, str := range strs {
		if len(str) > max { max = len(str) }
	}
	return max
}

type LenStruct interface {
	Len() int
}

type Str string
func (s Str)Len() int {
	return len(s)
}

type Strs []string
func (s Strs)Len() int {
	return getMaxLen(s)
}

func getMaxLen5(m LenStruct) int {
	return m.Len()
}

type LenRes struct {
	X int
	Y int
	Sum int
}

func main() {
	m := Strs{"sda","a"}
	fmt.Println(getMaxLen5(m))
}
