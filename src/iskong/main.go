package main

import (
	"fmt"
	"strings"
)

func iskong(l []interface{}) bool {
	if len(l) == 0 {
		return true
	}
	return false
}

func Name() []string {
	name := "sdfsdfsdf  1234"
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	nameSplit := strings.Split(name, "/")
	return nameSplit
}

func InSlice(i string,s []string) bool {
	for _, ss := range s {
		if i == ss {
			return true
		}
	}
	return false
}

func t1() {
	//s := []string{"sdfsdf", "sdfsdffff"}
	fmt.Println(Name())
	fmt.Println(InSlice("sdfs", Name()))
}

func main() {
	t1()
}
