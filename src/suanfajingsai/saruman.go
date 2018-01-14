package main

import "fmt"

var (
	N = 6
	R = 10
	X = []int{1, 7, 15, 20, 30, 50}
)

func saruman() int {
	count := 0
	i := 0
	var s, p int
	for i < N {
		s = X[i]
		for i < N && X[i] <= s + R {
			i++
		}
		p = X[i-1]
		for i < N && X[i] <= p + R {
			i++
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(saruman())
}
