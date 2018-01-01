package main

import "fmt"

var data []int

func init() {
	data = []int{3,1,5,7,2,8,4,9,6}
}

func insertSort() {
	l := len(data)
	for i:=1; i<l; i++ {
		if data[i] < data[i-1] {
			guard := data[i]
			j := i-1
			data[i] = data[i-1]
			for j>=0 && guard < data[j] {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = guard
		}
	}
}

func bubbleSort() {
	l := len(data)
	for i:=0; i<l; i++ {
		for j:=i; j<l; j++{
			if data[j] < data[i] {
				data[j], data[i] = data[i], data[j]
			}
		}
	}
}


func main() {
	//insertSort()
	bubbleSort()
	fmt.Println(data)
}
