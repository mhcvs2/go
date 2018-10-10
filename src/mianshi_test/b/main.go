package main

import "fmt"

type student struct {
	Name string
	Age int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name:"a", Age: 1},
		{Name:"b", Age: 2},
		{Name:"c", Age: 3},
	}

	//错
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
	fmt.Println("----------------------------")
	for i:=0; i<len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}

func main() {
	pase_student()
}

//a => c
//b => c
//c => c
//----------------------------
//a => a
//b => b
//c => c