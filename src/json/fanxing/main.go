package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code int
	Data interface{}
}

func t1() {
	r1 := Result{}
	r1.Code = 0
	r1.Data = []string{"asd","asdasd"}
	j,_ := json.Marshal(r1)
	fmt.Println(j)
	fmt.Println(string(j))
	r2 := Result{}
	json.Unmarshal(j, &r2)
	fmt.Println(r2)
	fmt.Println(r2.Code)
	for _, s := range r2.Data.([]string) {
		fmt.Println(s)
	}
}

func main() {
	t1()
}
