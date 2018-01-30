package main

import (
	"fmt"
	"os"
)

func t1() {
	fmt.Fprintf(os.Stdout, "haha%s\n", "ddd")
}
//hahaddd

func main() {
	t1()
}
