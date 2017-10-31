package main

import (
"fmt"
"sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
		close(out)
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
		close(out)
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 为每一个输入channel cs 创建一个 goroutine output
	// output 将数据从 c 拷贝到 out，直到 c 关闭，然后 调用 wg.Done
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// 启动一个 goroutine，用于所有 output goroutine结束时，关闭 out
	// 该goroutine 必须在 wg.Add 之后启动
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func t3() {

	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3)

	// 启动两个 sq 实例，即两个goroutines处理 channel "in" 的数据
	c1 := sq(done, in)
	c2 := sq(done, in)

	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9
	//fmt.Println(<-out)
}

func main() {
	t3()
}
