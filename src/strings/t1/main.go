package main

import (
	"io/ioutil"
	"fmt"
	"github.com/goinaction/code/chapter3/words"
)

func main() {
	filename := "/root/github/go/src/strings/t1/words.txt"

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Opening file error:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
