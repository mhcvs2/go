package main

import (
	"strings"
	"fmt"
)

func main()  {
	t1()
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