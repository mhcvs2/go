package main

import (
	"fmt"
	"reflect"
)

func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first
	for _, x := range rest {
		switch x:= x.(type) {
		case int:
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum = x
			}
		case string:
			if x < minimum.(string) {
				minimum = x
			}
		}
	}
	return minimum
}

func testMinimum()  {
	i := Minimum(4,3,8,2,9).(int)
	fmt.Printf("%T %v\n", i, i)
	f := Minimum(9.4, -5.4, 3.8, 17.0, -3.1, 0.0).(float64)
	fmt.Printf("%T %v\n", f, f)
	s := Minimum("K", "X", "B", "C", "CC", "CA", "D", "M").(string)
	fmt.Printf("%T %q\n", s, s)
	/*
	int 2
	float64 -5.4
	string "B"
	*/
}
//==================================================================================

func Index(xs interface{}, x interface{}) int {
	switch slice := xs.(type) {
	case []int:
		for i, y := range slice {
			if y == x.(int){
				return i
			}
		}
	case []string:
		for i, y := range slice {
			if y == x.(string){
				return i
			}
		}
	}
	return -1
}

func testIndex()  {
	xs := []int{2,4,6,8}
	fmt.Println("5 @", Index(xs, 5), "6 @", Index(xs, 6))
	ys := []string{"C", "B", "K", "A"}
	fmt.Println("Z @", Index(ys, "Z"), "A @", Index(ys, "A"))
	/*
	5 @ -1 6 @ 2
	Z @ -1 A @ 3
	*/
}
//=========================================================================================

func IndexReflectX(xs interface{}, x interface{}) int {
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			switch y := slice.Index(i).Interface().(type) {
			case int:
				if y == x.(int) {
					return i
				}
			case string:
				if y == x.(string) {
					return i
				}
			}
		}
	}
	return -1
}

func testIndexReflectX()  {
	xs := []int{2,4,6,8}
	fmt.Println("5 @", IndexReflectX(xs, 5), "6 @", IndexReflectX(xs, 6))
	ys := []string{"C", "B", "K", "A"}
	fmt.Println("Z @", IndexReflectX(ys, "Z"), "A @", IndexReflectX(ys, "A"))
	/*
	5 @ -1 6 @ 2
	Z @ -1 A @ 3
	*/
}
//=========================================================================================
func IndexReflect(xs interface{}, x interface{}) int {
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if reflect.DeepEqual(x, slice.Index(i).Interface()) {
				return i
			}
		}
	}
	return -1
}

func testIndexReflect()  {
	xs := []int{2,4,6,8}
	fmt.Println("5 @", IndexReflect(xs, 5), "6 @", IndexReflect(xs, 6))
	ys := []string{"C", "B", "K", "A"}
	fmt.Println("Z @", IndexReflect(ys, "Z"), "A @", IndexReflect(ys, "A"))
	/*
	5 @ -1 6 @ 2
	Z @ -1 A @ 3
	*/
}

//==================================================================================

type Slicer interface {
	EqualTo(i int, x interface{}) bool
	Len() int
}

//-----
type IntSlice []int

func (slice IntSlice) EqualTo(i int, x interface{}) bool {
	return slice[i] == x.(int)
}

func (slice IntSlice) Len() int {
	return len(slice)
}

//------
type StringSlice []string

func (slice StringSlice) EqualTo(i int, x interface{}) bool {
	return slice[i] == x.(string)
}

func (slice StringSlice) Len() int {
	return len(slice)
}

//-----
func IndexSlicer(slice Slicer, x interface{}) int {
	for i := 0; i < slice.Len(); i++ {
		if slice.EqualTo(i, x) {
			return i
		}
	}
	return -1
}

//-----
func IntIndexSlicer(ints []int, x int) int {
	return IndexSlicer(IntSlice(ints), x)
}

func StringIndexSlicer(strings []string, x string) int {
	return IndexSlicer(StringSlice(strings), x)
}

//------
func testIndexSlicer()  {
	xs := []int{2,4,6,8}
	fmt.Println("5 @", IntIndexSlicer(xs, 5), "6 @", IntIndexSlicer(xs, 6))
	ys := []string{"C", "B", "K", "A"}
	fmt.Println("Z @", StringIndexSlicer(ys, "Z"), "A @", StringIndexSlicer(ys, "A"))
	/*
	5 @ -1 6 @ 2
	Z @ -1 A @ 3
	*/
}

//==================================================================================
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

//==================================================================================
func main() {
	testSliceIndex()
}
