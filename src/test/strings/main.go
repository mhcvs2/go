package main

import (
	"strings"
	"fmt"
	"unicode"
)

func main()  {
	t3()
}

func t1()  {
	line2 := "  "
	res2 := strings.TrimSpace(line2)
	fmt.Println(res2)

	line := "aa ddf fgg"
	res := strings.TrimSpace(line)
	fmt.Println(res)

	line3 := "    aa ddf fgg    "
	res3 := strings.TrimSpace(line3)
	fmt.Println(res3)

	/*
	out:

	aa ddf fgg
	aa ddf fgg

	*/
}

func t2()  {
	line := "    aa ddf fgg    "
	for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
		fmt.Println(word)
	}
	/*
	aa
	ddf
	fgg
	*/
}

func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char)}
	return strings.FieldsFunc(s, notALetter)
}

func t3()  {
	fmt.Println(strings.Repeat("hi ", 5))
	// hi hi hi hi hi
}
