package main

import (
	"fmt"
	"os"
	gosignal "os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM}
	gosignal.Notify(c, signals...)

	//go func() {
	//	for sig := range c {
	//		fmt.Println(sig)
	//	}
	//}()
	//for s := range c {
	//	fmt.Println(s)
	//}
	<- c
	fmt.Println("quit...")
}
