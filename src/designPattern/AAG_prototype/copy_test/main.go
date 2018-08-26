package main

import "fmt"

type Thing struct {
	List []string
}

func (t *Thing) Clone() *Thing {
	newThing := *t
	return &newThing
}

func main() {
	thing := new(Thing)
	cloneThing := thing.Clone()
	thing.List = append(thing.List, "zhangsan")
	cloneThing.List = append(cloneThing.List, "lisi")
	fmt.Println(cloneThing.List)
}
//[lisi]
