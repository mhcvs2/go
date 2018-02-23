package main

import (
	"time"
	"fmt"
	"math/rand"
)

func t1() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<10; i++ {
		fmt.Println(r.Intn(2))
	}
}

func main() {
	t1()
}