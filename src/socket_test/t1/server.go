package main

import (
	"fmt"
	"net"
	"os"
	gosignal "os/signal"
	"syscall"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT}
	gosignal.Notify(c, signals...)

	go func() {
		fmt.Println("wait to quit...")
		<- c
		fmt.Println("quit...")
		os.Exit(0)
	}()

	listener, err := net.Listen("tcp", ":18000")
	handleErr(err)
	defer listener.Close()
	fmt.Println("listen: ", listener.Addr())

	conn, err := listener.Accept()
	handleErr(err)
	defer conn.Close()
	fmt.Println("local addr: ", conn.LocalAddr())
	fmt.Println("remote addr", conn.RemoteAddr())



	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		handleErr(err)
		fmt.Println("server get: ", string(buf[:n]))
	}
}

