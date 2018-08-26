package main

import (
	"fmt"
)

type Text struct {
	Name string
}

type Thing struct {
	n1 Text
	n2 *Text
}

func (t *Thing) Clone() *Thing {
	newThing := *t
	return &newThing
}

func main() {
	t1 := Text{Name:"t1"}
	t2 := &Text{Name:"t2"}
	thing := &Thing{n1:t1,n2:t2}
	cloneThing := thing.Clone()

	thing.n1.Name = "h1"
	thing.n2.Name = "h2"

	fmt.Println(cloneThing.n1.Name)
	fmt.Println(cloneThing.n2.Name)
}
//t1
//h2
