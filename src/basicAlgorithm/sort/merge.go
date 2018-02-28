package main

import "fmt"

var count = 0

func MergeSort(arr []int) {
	l := len(arr)
	tmp := make([]int, l)
	mergeSort(arr, 0, l-1, tmp)
}

func mergeSort(arr []int, left, right int, tmp []int) {
	if left < right{
		mid := (left + right) / 2
		mergeSort(arr, left, mid, tmp)
		mergeSort(arr, mid+1, right, tmp)
		merge(arr, left, mid, right, tmp)
	}
}

func merge(arr []int, left, mid, right int, tmp []int) {
	i := left
	j := mid + 1
	t := left
	for i <= mid && j <= right {
		count++
		if arr[i] < arr[j] {
			tmp[t] = arr[i]
			i++
		}else {
			tmp[t] = arr[j]
			j++
		}
		t++
	}
	for i <= mid {
		count++
		tmp[t] = arr[i]
		i++
		t++
	}
	for j <= right {
		count++
		tmp[t] = arr[j]
		j++
		t++
	}
	for k := left; k <= right; k++ {
		count++
		arr[k] = tmp[k]
	}
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	MergeSort(arr)
	fmt.Println(arr)
	fmt.Println(count)
}