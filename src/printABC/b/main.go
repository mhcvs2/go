package main

import (
	"fmt"
	"sync"
)

var (
	state = 0
	lock sync.Mutex
	wg sync.WaitGroup
)

func pa() {
	defer wg.Done()
	for i:=0; i<10; {
		lock.Lock()
		if state % 3 == 0 {
			fmt.Println("A")
			state++
			i++
		}
		lock.Unlock()
	}
}

func pb() {
	defer wg.Done()
	for i:=0; i<10;{
		lock.Lock()
		if state % 3 == 1 {
			fmt.Println("B")
			state++
			i++
		}
		lock.Unlock()
	}
}

func pc() {
	defer wg.Done()
	for i:=0; i<10; {
		lock.Lock()
		if state % 3 == 2 {
			fmt.Println("C")
			state++
			i++
		}
		lock.Unlock()
	}
}

func main() {
	go pa()
	go pb()
	go pc()

	wg.Add(3)
	wg.Wait()

}
