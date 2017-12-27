package main

import "fmt"

var data []int

func init() {
	data = []int{3,1,5,7,2,8,4,9,6}
}

func insertSortDown() {
	l := len(data)
	for i:=1; i<l; i++ {
		if data[i] > data[i-1] {
			guard := data[i]
			data[i] = data[i-1]
			j := i - 1
			for {
				if j < 0 || guard <= data[j] {
					break
				}
				data[j+1] = data[j]
				j--
			}
			data[j+1] = guard
		}
	}
}

func bubbleSortDown() {
	l := len(data)
	for i:=0; i<l; i++{
		for j:=i;j<l; j++{
			if data[j] > data[i] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

func main() {
	//insertSortDown()
	bubbleSortDown()
	fmt.Println(data)
}