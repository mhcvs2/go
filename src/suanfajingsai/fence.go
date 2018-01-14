package main

import "fmt"

var (
	N = 3
	L = []int{8, 5, 8}
)

func sort() {
	for i:=0; i< N; i++ {
		for j:=i; j < N; j++ {
			if L[j] < L[i] {
				tmp := L[j]
				L[j] = L[i]
				L[i] = tmp
			}
		}
	}
}



func fence() {
	sort()
	fmt.Println(L)
	L2 := make([]int, N)
	sum := L[0]
	for i:=1; i<N; i++ {
		sum += L[i]
		L2[i-1] = sum
	}
	fmt.Println(L2)
	res := 0
	for n := range L2 {
		res += L2[n]
	}
	fmt.Println(res)
}

func main() {
	fence()
}
