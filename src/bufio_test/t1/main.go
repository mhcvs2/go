package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func t1() {
	if f, err := os.Open("/root/github/go/src/bufio_test/t1/a.txt"); err != nil {
		fmt.Println("open error: ", err.Error())
	} else {
		buf := bufio.NewReader(f)
		for{
			var outputbuf bytes.Buffer
			out, end, err := buf.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("read error: ", err.Error())
				break
			}
			outputbuf.Write(out)
			fmt.Println(outputbuf.String())
			if end {
				break
			}
		}
	}
}
//sdfsadf
//sdfg
//hello
//kkkkk

func t2() {
	if f, err := os.Open("/root/github/go/src/bufio_test/t1/a.txt"); err != nil {
		fmt.Println("open error: ", err.Error())
	} else {
		var outputbuf bytes.Buffer
		for {
			tempOutput := make([]byte, 5)
			n, err := f.Read(tempOutput)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println("read error: ", err.Error())
				}
			}
			if n > 0 {
				outputbuf.Write(tempOutput[:n])
			}
		}
		fmt.Println(outputbuf.String())
	}
}
//sdfsadf
//sdfg
//hello
//kkkkk

func main() {
	t2()
}
