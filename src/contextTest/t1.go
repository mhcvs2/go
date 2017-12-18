package main

import (
	"context"
	"time"
	"fmt"
)

func t11() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout!")
	default:
		fmt.Println("default")
	}
}
//default

func t12() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer func() {
		fmt.Println("cancle")
		cancel()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout!")
	}
}
//Timeout!
//cancle

func t13() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("cancle")
		cancel()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout!")
	}
}
//cancle
//Timeout!

func main() {
	t12()
}
