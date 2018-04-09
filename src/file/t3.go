package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"fmt"
)

func Readlines(filePath string, handler func(line []byte) error) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		line = bytes.TrimRight(line, "\n\r")
		err = handler(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func t31() {
	err := Readlines("/root/github/go/src/chan/cgrep/testfiles/a.txt", func(line []byte) error {
		fmt.Println(string(line))
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	t31()
}