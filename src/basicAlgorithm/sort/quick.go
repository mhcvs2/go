package main

import (
	"fmt"
	"time"
	"math/rand"
)

// nlogn
func quickSort(arr []int, _left, _right int) {
	left := _left
	right := _right
	if left <= right {
		for left < right {
			tmp := arr[left]
			for left < right && arr[right] >= tmp {
				right--
			}
			arr[left] = arr[right]
			for left < right && arr[left] <= tmp {
				left++
			}
			arr[right] = arr[left]
			arr[left] = tmp
		}
		quickSort(arr, _left, left-1)
		quickSort(arr, right+1, _right)
	}
}

func quickSort2(arr []int, _left, _right int) {
	left := _left
	right := _right
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if left <= right {
		randIndex := r.Intn(right-left+1) + left
		tmp := arr[left]
		arr[left] = arr[randIndex]
		arr[randIndex] = tmp
		for left < right {
			tmp = arr[left]
			for left < right && arr[right] >= tmp {
				right--
			}
			arr[left] = arr[right]
			for left < right && arr[left] <= tmp {
				left++
			}
			arr[right] = arr[left]
			arr[left] = tmp
		}
		quickSort(arr, _left, left-1)
		quickSort(arr, right+1, _right)
	}
}

func main() {
	arr := []int{5, 10, 3, 1, 7, 2, 8}
	quickSort2(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
