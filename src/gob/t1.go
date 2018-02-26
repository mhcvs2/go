package main

import (
	"fmt"
	"encoding/gob"
	"bytes"
)

type MyFace interface {
	A()
}

type Cat struct {}
type Dog struct {}

func (c *Cat) A() {
	fmt.Println("Meow")
}

func (d *Dog) A() {
	fmt.Println("Woof")
}

func init() {
	gob.Register(&Cat{})
	gob.Register(&Dog{})
}

func main() {
	network := new(bytes.Buffer)
	enc := gob.NewEncoder(network)

	var inter MyFace
	inter = new(Cat)

	err := enc.Encode(&inter)
	if err != nil {
		panic(err)
	}

	inter = new(Dog)
	err = enc.Encode(&inter)
	if err != nil {
		panic(err)
	}

	dec := gob.NewDecoder(network)

	var get MyFace
	err = dec.Decode(&get)
	if err != nil {
		panic(err)
	}

	get.A()
}
