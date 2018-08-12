package main

import (
	"regexp"
	"fmt"
)


var text2 = "requests.nvidia.com/gpu-k20m"

func a2() {
	pattern := "cpu|memory|gpu.*"
	obj := regexp.MustCompile(pattern)
	fmt.Println(obj.FindString(text2))
}

func main() {
	a2()
}
