package main

import "fmt"

type Student struct {
	Id int
	Name string
}


var Data = make(map[string]*Student)

func t1() {
	stus := []Student{
		{Id:1, Name:"a"},
		{Id:2, Name:"b"},
		{Id:3, Name:"c"},
	}

	for _, stu := range stus {
		fmt.Println(stu.Id)
		Data[stu.Name] = &stu
	}

	for k, v := range Data {
		fmt.Printf("k: %s, v: %v", k, v)
	}
}
//1
//2
//3
//k: a, v: &{3 c}k: b, v: &{3 c}k: c, v: &{3 c}

func t2() {
	stus := []*Student{
		{Id:1, Name:"a"},
		{Id:2, Name:"b"},
		{Id:3, Name:"c"},
	}

	for _, stu := range stus {
		fmt.Println(stu.Id)
		Data[stu.Name] = stu
	}

	for k, v := range Data {
		fmt.Printf("k: %s, v: %v", k, v)
	}
}
//1
//2
//3
//k: a, v: &{1 a}k: b, v: &{2 b}k: c, v: &{3 c}

func main() {
	t2()
}