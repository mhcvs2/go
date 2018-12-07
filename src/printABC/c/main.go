package main
// 模拟java里的condition，python里的threading.event

import (
	"fmt"
	"sync"
)

var (
	a chan int
	b chan int
	c chan int
	state = 0
	wg sync.WaitGroup
)

func init() {
	a = make(chan int, 1)
	b = make(chan int, 1)
	c = make(chan int, 1)
}

func pa() {
	defer wg.Done()
	for i:=0; i<10; i++ {
		if state % 3 != 0 {
			<- a
		}
		fmt.Println("A")
		state++
		b <- 1    // 管道的缓冲大小是1，是防止这里阻塞
	}
}

func pb() {
	defer wg.Done()
	for i:=0; i<10; i++ {
		if state % 3 != 1 {
			<- b
		}
		fmt.Println("B")
		state++
		c <- 1
	}
}

func pc() {
	defer wg.Done()
	for i:=0; i<10; i++ {
		if state % 3 != 2 {
			<- c
		}
		fmt.Println("C")
		state++
		a <- 1
	}
}

func main() {
	go pa()
	go pb()
	go pc()

	wg.Add(3)
	wg.Wait()

}