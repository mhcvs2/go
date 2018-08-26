package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
)

//字节流拷贝，无引用共享问题
func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

type Text struct {
	Name string
}

type Thing struct {
	N1 Text
	N2 *Text
}

func (t *Thing) Clone() *Thing {
	newThing := new(Thing)
	err := deepCopy(newThing, t)
	if err != nil {
		panic(err)
	}
	return newThing
}

func main() {
	t1 := Text{Name:"t1"}
	t2 := &Text{Name:"t2"}
	thing := &Thing{N1:t1,N2:t2}
	cloneThing := thing.Clone()

	thing.N1.Name = "h1"
	thing.N2.Name = "h2"

	fmt.Println(cloneThing.N1.Name)
	fmt.Println(cloneThing.N2.Name)
}
//t1
//t2
