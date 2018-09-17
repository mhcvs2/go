package main

import (
	"io"
	"crypto/rand"
	"encoding/hex"
)

func RandomString() string {
	id := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, id); err != nil {
		panic(err)
	}
	return hex.EncodeToString(id)
}


func t1() {
	println(RandomString())
}

func main() {
	t1()
}
