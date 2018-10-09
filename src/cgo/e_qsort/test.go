package main

import (
	"cgo/e_qsort/qsort"
	"fmt"
)

func main() {
	values := []int64{42, 9, 101, 95, 27, 25}
	qsort.Slice(values, func(a, b int) bool {
		return values[a] < values[b]
	})
	fmt.Println(values)
}
