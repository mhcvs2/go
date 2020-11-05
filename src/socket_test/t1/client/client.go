package main

import (
	"fmt"
	"net"
)

func hErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:18000")
	hErr(err)
	defer conn.Close()
	fmt.Println("local addr: ", conn.LocalAddr())
	fmt.Println("remote addr", conn.RemoteAddr())

	_, err = conn.Write([]byte("Are u ready?"))
	hErr(err)
}
