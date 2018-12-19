package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	in chan int
	a chan int
	b chan int
	c chan int
	done chan struct{}
	//wg sync.WaitGroup
	lock sync.Mutex
)

func init() {
	in = make(chan int, 1)
	a = make(chan int, 1)
	b = make(chan int, 1)
	c = make(chan int, 1)
	done = make(chan struct{})
}

func printA() {
	for{
		select {
		case num := <-a:
			if num % 3 == 0 {
				fmt.Println("A")
			} else {
				b <- num
			}
		case <-done:
			break
		}
	}
}

func printB() {
	for{
		select {
		case num := <-b:
			if num % 3 == 1 {
				fmt.Println("B")
			} else {
				c <- num
			}
		case <-done:
			break
		}
	}
}

func printC() {
	for{
		select {
		case num := <-c:
			if num % 3 == 2 {
				fmt.Println("C")
			}
		case <-done:
			break
		}
	}
}

func main() {
	go printA()
	go printB()
	go printC()

	for i:=0; i<30; i++ {
		a <- i
	}
	time.Sleep(time.Second * 2)
	close(done)

}
