package main

import (
	"fmt"
	"time"
)

func t1() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now())
	expirationTime := <-timer.C
	fmt.Println(time.Now())
	fmt.Println(expirationTime)
	fmt.Println(timer.Stop())

}
//2019-01-30 16:48:10.888045282 +0800 CST m=+0.000159159
//2019-01-30 16:48:12.888132393 +0800 CST m=+2.000246246
//2019-01-30 16:48:12.888117019 +0800 CST m=+2.000230892
//false

func t2() {
	intChan := make(chan int, 1)
	go func() {
		for i:=0; i<5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()
	timeout := time.Millisecond * 500
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("end")
				return
			}
			fmt.Println("received", e)
		case <-timer.C:
			fmt.Println("timeout")
		}
	}
}

func main() {
	t2()
}
