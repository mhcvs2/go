package main

import "fmt"

func t11() {
	FOR1:
		for i:=0; i<100; i++ {
			fmt.Println("i-------------")
			fmt.Println(i)
			for j :=200; j < 300; j++ {
				fmt.Println("j--------------")
				fmt.Println(j)
				if j > 205 {
					break FOR1
				}
			}
		}
}

func main() {
	t11()
}