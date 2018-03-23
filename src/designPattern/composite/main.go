package main

import (
	"fmt"
	"reflect"
)

type ICorp interface {
	GetInfo()
}

type Corp struct {
	name string
	salary int
}

func (c *Corp)GetInfo() {
	fmt.Println("name: ", c.name)
	fmt.Println("salary", c.salary)
	fmt.Println()
}

type Leaf struct {
	Corp
}

func NewLeaf(name string, salary int) *Leaf {
	return &Leaf{Corp:Corp{name:name, salary:salary}}
}

type Branch struct {
	Corp
	sub []ICorp
}

func NewBranch(name string, salary int) *Branch {
	return &Branch{
		Corp:Corp{name:name, salary:salary},
		sub:[]ICorp{},
	}
}

func (b *Branch)AddSub(corp ICorp) {
	b.sub = append(b.sub, corp)
}

func (b *Branch) GetSub() []ICorp {
	return b.sub
}

/////////////////////////////////////////////////

func GetTreeInfo(root *Branch) {
	for _, b := range root.GetSub() {
		vb := reflect.ValueOf(b)
		if vb.Type() == reflect.ValueOf(new(Leaf)).Type() {
			b.GetInfo()
		} else {
			b.GetInfo()
			GetTreeInfo(b.(*Branch))
		}
	}
}

func main() {
	ceo := NewBranch("ceo", 1000000)
	l1 := NewBranch("leader1", 10000)
	l2 := NewBranch("leader2", 10001)

	y1 := NewLeaf("y1", 1)
	y2 := NewLeaf("y2", 2)
	y3 := NewLeaf("y3", 3)

	ceo.AddSub(l1)
	ceo.AddSub(l2)
	ceo.AddSub(y1)

	l1.AddSub(y2)
	l2.AddSub(y3)

	ceo.GetInfo()
	GetTreeInfo(ceo)

}

//name:  ceo
//salary 1000000
//
//name:  leader1
//salary 10000
//
//name:  y2
//salary 2
//
//name:  leader2
//salary 10001
//
//name:  y3
//salary 3
//
//name:  y1
//salary 1