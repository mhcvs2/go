package main

import (
	"fmt"
	"math"
)

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func testSliceIndex()  {
	xs := []int{2,4,6,8}
	ys := []string{"C", "B", "K", "A"}
	fmt.Println(
		SliceIndex(len(xs), func(i int) bool { return xs[i] == 5}),
		SliceIndex(len(xs), func(i int) bool { return xs[i] == 6}),
		SliceIndex(len(ys), func(i int) bool { return ys[i] == "Z"}),
		SliceIndex(len(ys), func(i int) bool { return ys[i] == "A"}))
	/*
	-1 2 -1 3
	*/
}

func test2()  {
	i := SliceIndex(math.MaxInt32, func(i int) bool {
		return i > 0 && i%27 ==0 && i%51 == 0
	})
	fmt.Println(i)
	//459    能被27和51整除的最小整数
}

//===========================================================
func IntFilter(slice []int, predicate func(int) bool) []int {
	filters := make([]int, 0, len(slice))
	for i:= 0; i < len(slice); i++ {
		if predicate(slice[i]) {
			filters = append(filters, slice[i])
		}
	}
	return filters
}

func testIntFilter() {
	readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even := IntFilter(readings, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(even)
	//[4 2 8 18 -6]       过滤掉奇数
}

//=============================================================
func Filter(limit int, predicate func(int) bool, appender func(int)) {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			appender(i)
		}
	}
}

func testInt() {
	readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even := make([]int, 0, len(readings))
	Filter(len(readings), func(i int) bool {
		return i%2 == 0
	}, func(i int) {
		even = append(even, readings[i])
	})
	fmt.Println(even)
	// [4 2 8 -11 18]
}

func testString() {
	readings := []string{"X15", "T14", "X23", "A41", "L19", "X57", "A63"}
	var even []string
	Filter(len(readings), func(i int) bool {
		return readings[i][0] == 'X'
	}, func(i int) {
		even = append(even, readings[i])
	})
	fmt.Println(even)
	// [X15 X23 X57]
}

func test3()  {
	var product int64 = 1
	Filter(26, func(i int) bool {
		return i%2 != 0
	}, func(i int) {
		product *= int64(i)
	})
	fmt.Println(product)
	// 7905853580625 计算[1, 25]内所有奇数的乘积
}

//=========================================================
func main()  {
	test3()
}
