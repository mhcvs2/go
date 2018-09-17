package main

import (
	"strings"
	"fmt"
)

func t1() {
	driver := "srcb/cephfs"
	driver = strings.Replace(driver, "/", "~", 1)
	fmt.Println(driver)
}

func main() {
	t1()
}
