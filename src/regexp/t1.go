package main

import (
	"regexp"
	"fmt"
)

var text = `Hello 世界！123 Go.`

func t11() {
	reg := regexp.MustCompile(`[a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//["ello" "o"]
}

func t12() {
	reg := regexp.MustCompile(`[a-z]+`)
	reg2 := reg.Copy()
	fmt.Printf("%q\n", reg2.FindAllString(text, -1))
	//["ello" "o"]
}

func t13() {
	reg := regexp.MustCompile(`[a-z]+`)
	m := reg.MatchString(text)
	fmt.Println(m)
	//true
}

func t14() {
	reg := regexp.MustCompile(`[a-z]+`)
	m := reg.ReplaceAllString(text, "mhc")
	fmt.Println(m)
	//Hmhc 世界！123 Gmhc.
}

func main() {
	t14()
}
