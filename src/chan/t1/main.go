package main

import (
	"fmt"
	"time"
)

func print1(ch chan int) {
	fmt.Println("Hello World.")
	ch <- 1
}

func print2(ch chan int) {
	ch <- 1
	fmt.Println("Hello World.")
}

func t1() {
	chs := make([]chan int,10)
	for i:=0; i<10; i++ {
		chs[i] = make(chan int)
		go print2(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func t2() {
	s := []int{7,2,8,-9,4,0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func t3() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i< 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c{
		fmt.Println(i)
	}

	fmt.Println("Finished")
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func t4() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i:=0; i<10;i++{
			fmt.Println(<-c)
		}
		quit <- 0

	}()
	fibonacci(c,quit)
}

func t5() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
}

func worker(done chan bool) {
	time.Sleep(time.Second)

	done <- true
}

func t6() {
	done := make(chan bool, 1)
	go worker(done)

	<- done
}

func main() {
	t6()
}
