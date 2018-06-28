package main

import (
	"sync/atomic"
	"time"
	"fmt"
	"sync"
	"github.com/panjf2000/ants"
)

var sum int32

func myFunc(i interface{}) error {
	n := i.(int)
	atomic.AddInt32(&sum, int32(n))
	return nil
}

func demoFunc() error {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World")
	return nil
}

func t1() {
	runTimes := 1000
	var wg sync.WaitGroup
	for i := 0; i<runTimes; i++ {
		wg.Add(1)
		ants.Submit(func() error {
			demoFunc()
			wg.Done()
			return nil
		})
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}

func t2() {
	runTimes := 1000
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) error {
		myFunc(i)
		wg.Done()
		return nil
	})
	for i:=0; i<runTimes; i++ {
		wg.Add(1)
		p.Serve(i)
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks. result is %d\n", sum)
}

func main() {
	t2()
}
