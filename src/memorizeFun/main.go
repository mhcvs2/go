package main

import (
	"fmt"
	"bytes"
	"strings"
)

type memorizeFunction func(int, ...int) interface{}

func Memorize(function memorizeFunction) memorizeFunction {
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprint(",%d", i)
		}
		if value, found := cache[key]; found {
			return value
		}
		value := function(x, xs...)
		cache[key] = value
		return value
	}
}

func testFibonacci()  {
	var Fibonacci memorizeFunction
	Fibonacci = Memorize(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}
		return Fibonacci(x-1).(int) + Fibonacci(x-2).(int)
	})
	fmt.Println("Fibonacci(45) = ", Fibonacci(45).(int))
	//Fibonacci(45) =  1134903170
}

var RomanForDEcimal memorizeFunction
func init()  {
	//可以将Memorize的使用放到这里，全局访问

	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	RomanForDEcimal = Memorize(func(x int, xs ...int) interface{} {
		if x < 0 || x > 3999 {
			panic("RomanForDEcimal() only handles integers [0, 3999]")
		}
		var buffer bytes.Buffer
		for i, decimal := range decimals {
			remainder := x / decimal
			x %= decimal
			if remainder > 0 {
				buffer.WriteString(strings.Repeat(romans[i], remainder))
			}
		}
		return buffer.String()
	})
}

func test()  {
	fmt.Println(RomanForDEcimal(1))
	fmt.Println(RomanForDEcimal(3))
	fmt.Println(RomanForDEcimal(5))
	fmt.Println(RomanForDEcimal(7))
	fmt.Println(RomanForDEcimal(666))

	/*
	I
	III
	V
	VII
	DCLXVI
	*/
}

func main()  {
	test()
}
