package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock sync.Mutex
	wg sync.WaitGroup
)

func a() {
	lock.Lock()
	defer wg.Done()
	defer lock.Unlock()
	fmt.Println("aaaaaaaaaaaaaa")
	time.Sleep(time.Second * 5)
}

func b() {
	lock.Lock()
	defer wg.Done()
	defer lock.Unlock()
	fmt.Println("bbbbbbbb")
	time.Sleep(time.Second * 5)
}

func main() {
	go a()
	go b()
	wg.Add(2)
	wg.Wait()

}