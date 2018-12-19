package main

import (
	"fmt"
	"sync"
)

//模拟信号量

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
	a <- 1
}

func pa() {
	defer wg.Done()
	for i:=0; i<10; i++ {
		<- a //获取信号量
		fmt.Println("A")
		b <- 1 //释放b信号量
	}
}

func pb() {
	defer wg.Done()
	for i:=0; i<10; i++ {
		<- b
		fmt.Println("B")
		c <- 1
	}
}

func pc() {
	defer wg.Done()
	for i:=0; i<10; i++ {
		<- c
		fmt.Println("C")
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
