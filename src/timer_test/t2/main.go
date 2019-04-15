package main

import (
	"fmt"
	"time"
)

func t1() {
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:

			}
		}
		fmt.Println("sender end")
	}()
	var sum int
	for e:= range intChan {
		fmt.Println("received ", e)
		sum += e
		if sum > 10 {
			fmt.Println("got ", sum)
			break
		}
	}
	fmt.Println("receiver end")
}

func main() {
	t1()
}
