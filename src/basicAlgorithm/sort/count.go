package main

import "fmt"

func countSort(arr []int) {
	c := make([]int, 1000)
	l := len(arr)
	for _, v := range arr {
		c[v] += 1
	}
	for i:=1; i<1000; i++ {
		c[i] += c[i-1]
	}
	res := make([]int, l)
	for i:=0; i<l; i++ {
		res[c[arr[i]]-1] = arr[i]
		c[arr[i]] -= 1
	}
	for i, v := range res {
		arr[i] = v
	}
}

func main() {
	arr := []int{4, 5, 3, 1, 7, 2, 8}
	countSort(arr)
	fmt.Println(arr)
}