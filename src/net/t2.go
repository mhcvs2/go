package main

import (
	"net"
	"fmt"
	"os"
	"io/ioutil"
	"time"
)

func t21() {
	port, err := net.LookupPort("tcp", "telnet")

	if err != nil {
		fmt.Fprintf(os.Stderr, "未找到指定服务")
		return
	}
	fmt.Fprintf(os.Stdout, "telnet port: %d", port)
}
//telnet port: 23

func t22() {
	// TCPAddr 包涵IP 和 Port

	// 将一个host地址转换为TCPAddr。host=ip:port
	pTCPAddr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "www.baidu.com:80 IP: %s PORT: %d", pTCPAddr.IP.String(), pTCPAddr.Port)
}

func t23() {
	url := "www.baidu.com:80"

	pRemoteTCPAddr, err := net.ResolveTCPAddr("tcp4", url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	// pLocalTCPAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7070")

	pTCPConn, err := net.DialTCP("tcp", nil/*pLocalTCPAddr*/, pRemoteTCPAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}


	n, err := pTCPConn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}
	defer pTCPConn.Close()

	fmt.Fprintf(os.Stdout, "writed: %d\n", n)

	buf, err := ioutil.ReadAll(pTCPConn)
	r := len(buf)
	fmt.Fprintf(os.Stdout, string(buf[:r]))
	fmt.Fprintf(os.Stdout, "readed: %d\n", r)
}

func t24() {
	pTCPAddr, error := net.ResolveTCPAddr("tcp4", ":7070")
	if error != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", error.Error())
		return
	}
	pTCPListener, error := net.ListenTCP("tcp4", pTCPAddr)
	if error != nil {
		fmt.Fprintf(os.Stdout, "Error: %s", error.Error())
		return
	}
	defer pTCPListener.Close()

	for {
		pTCPConn, error := pTCPListener.AcceptTCP()
		if error != nil {
			fmt.Fprintf(os.Stdout, "Error: %s", error.Error())
			continue
		}
		go connHandler(pTCPConn)
	}
}

func connHandler(conn *net.TCPConn) {
	defer conn.Close()
	now := time.Now()
	conn.Write([]byte(now.String() + "\n"))
}

func main() {
	t24()
}