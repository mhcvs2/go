package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func t1() {
	// 设置流水线
	c := gen(2, 3)
	out := sq(c)

	// 消费输出结果
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func t2() {
	// 设置流水线并消费输出结果
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 为每一个输入channel cs 创建一个 goroutine output
	// output 将数据从 c 拷贝到 out，直到 c 关闭，然后 调用 wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
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
	in := gen(2, 3)

	// 启动两个 sq 实例，即两个goroutines处理 channel "in" 的数据
	c1 := sq(in)
	c2 := sq(in)

	// merge 函数将 channel c1 和 c2 合并到一起，这段代码会消费 merge 的结果
	for n := range merge(c1, c2) {
		fmt.Println(n) // 打印 4 9, 或 9 4
	}
}

func main() {
	t3()
}