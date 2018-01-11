package main

import "fmt"

//最少用多少硬币

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(v, c []int, a int) int {
	l := len(v)
	count := 0
	aCount := 0
	for i:=l-1; i>=0; i-- {
		if a == 0 { break }
		aCount = min(a/v[i], c[i])
		count += aCount
		a -= v[i] * aCount
	}
	if a != 0 {
		return -1
	}
	return count
}

func main() {
	v := []int{1, 5, 10, 50, 100, 500}
	c := []int{2, 2, 2, 2, 2, 2}
	a := 1010
	fmt.Println(solve(v, c, a))
}
