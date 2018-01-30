package main

import "fmt"

func erfenSearch(datas []int, data int) int {
	l := len(datas)
	if l == 0 {
		return -1
	}
	low := 0
	top := l-1
	pos := -1
	for low <= top {
		pos = (low + top)/2
		if datas[pos] == data {
			return pos
		}
		if datas[pos] > data {
			top = pos-1
		} else {
			low = pos + 1
		}
	}
	return -1
}

func erfenSearch2(datas []int, data, low, top int) int {
	if low > top || len(datas) == 0 {
		return -1
	}
	pos := (low + top)/2
	if datas[pos] == data {
		return pos
	}
	if datas[pos] > data {
		return erfenSearch2(datas, data, low, pos-1)
	} else {
		return erfenSearch2(datas, data, pos+1, top)
	}
	return -1
}

func erfenSearch3(datas []int, data int, left bool) int {
	l := len(datas)
	if l == 0 {
		return -1
	}
	low := 0
	top := l-1
	pos := -1
	last := -1
	for low <= top {
		pos = (low + top)/2
		if datas[pos] == data {
			last = pos
			if left {
				top = pos - 1
			} else {
				low = pos + 1
			}
		} else if datas[pos] > data {
			top = pos-1
		} else {
			low = pos + 1
		}
	}
	return last
}

func main() {
	//fmt.Println(erfenSearch([]int{1,2,3,4,5,6,7,8,9}, 10))
	//fmt.Println(erfenSearch2([]int{1,2,3,4,5,6,7,8,9}, 9, 0, 8))
	fmt.Println(erfenSearch3([]int{1,2,3,4,5,5,5,5,5,5,6,7,8,9}, 5, true))
	fmt.Println(erfenSearch3([]int{1,2,3,4,5,5,5,5,5,5,6,7,8,9}, 5, false))
}
