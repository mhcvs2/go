package main

import "fmt"

type Father struct {
}

func (f *Father) A() {
	fmt.Println("Father A...")
}

func (f *Father) B() {
	fmt.Println("Father B will Call A")
	f.A()
}

type Son struct {
	Father
}

func (s *Son) A() {
	fmt.Println("Son A...")
}

func main() {
	s := new(Son)
	s.B()
}

//Father B will Call A
//Father A...