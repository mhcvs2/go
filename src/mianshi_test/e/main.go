package main

import (
	"runtime"
	"fmt"
)

func t1() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <- int_chan:
		fmt.Println(value)
	case value := <- string_chan:
		panic(value)
	}
}
//select会随机选择一个可用通用做收发操作。 所以代码是有肯触发异常，也有可能不会。
// 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：
//select 中只要有一个case能return，则立刻执行。
//当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
//如果没有一个case能return则可以执行”default”块。

func t2() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int)
	string_chan := make(chan string)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <- int_chan:
		fmt.Println(value)
	case value := <- string_chan:
		panic(value)
	}
}
//fatal error: all goroutines are asleep - deadlock!
//因为无缓冲会堵塞，没有活动的goroutine可以取

func t3() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int)
	string_chan := make(chan string)
	go func() {
		for {
			select {
			case value := <- int_chan:
				fmt.Println(value)
			case value := <- string_chan:
				fmt.Println(value)
			default:
				break
			}
		}
	}()
	int_chan <- 1
	string_chan <- "hello"
}
//1
//hello


func main() {
	t3()
}
