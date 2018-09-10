package main

import (
	"time"
	"fmt"
)

/*
close channel 之后select可以返回
*/

func t1() {
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 5)
		close(done)
	}()

	select {
	case <-time.After(time.Second * 4):
		fmt.Println("timeout")
	case <-done:
		fmt.Println("done")

	}
}

func main() {
	t1()
}