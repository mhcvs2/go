package main

import "fmt"

func test(s ...string) {
	for _, ss := range s {
		fmt.Println(ss)
	}
}

func t1() {
	test("aaa","dddd")
}

func t2() {
	paras := []string{"sadf",
	                  "sdfdfssss"}
	test(paras)

}


func main() {
	t1()
}