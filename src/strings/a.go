package main

import (
	"fmt"
	"strings"
)

func a1() {
	name := "sfdsdf.ddddd.fgggfg"
	fmt.Println(strings.Split(name, "."))
	fmt.Println(strings.SplitAfter(name, "."))
	fmt.Println(strings.SplitAfterN(name, ".", 2))
}
//[sfdsdf ddddd fgggfg]
//[sfdsdf. ddddd. fgggfg]
//[sfdsdf. ddddd.fgggfg]


func a2() {
	name := "sfdsdf.ddddd.fgggf	hahaha*lalalal"
	fmt.Println(strings.FieldsFunc(name, func(r rune) bool {
		switch r {
		case '\t', '*', '.':
			return true
		}
		return false
	}))
}
//[sfdsdf ddddd fgggf hahaha lalalal]

func main() {
	a2()
}
