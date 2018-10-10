package main

import "fmt"

//错
//func t1() {
//	for i:=0; i<10; i++ {
//		loop:
//			println(i)
//	}
//	goto loop
//}

func t2(i int) {
	if i % 2 == 0 {
		goto one
	} else {
		goto two
	}
	one:
		fmt.Println("ou")
	two:
		fmt.Println("ji")
}

func main() {
	t2(3)
}
