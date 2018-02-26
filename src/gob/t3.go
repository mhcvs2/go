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

func (c *Cat) A() {
	fmt.Println("Meow")
}

func init() {
	gob.Register(&Cat{})
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

	dec := gob.NewDecoder(bytes.NewBufferString(network.String()))

	var get MyFace
	err = dec.Decode(&get)
	if err != nil {
		panic(err)
	}

	get.A()
}
