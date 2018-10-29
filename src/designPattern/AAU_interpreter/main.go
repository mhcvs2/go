package main

import (
	"fmt"
	"designPattern/AAU_interpreter/calculator"
)

func getExpStr() string {
	fmt.Print("请输入表达式:")
	var expStr string
	fmt.Scanln(&expStr)
	return expStr
}

func getValue(expStr string)  map[string]int {
	v := make(map[string]int)
	for _, ch := range []byte(expStr){
		if ch != '+' && ch != '-'{
			if _, ok := v[string(ch)]; !ok {
				var in int
				fmt.Printf("输入%s的值: ", string(ch))
				fmt.Scanln(&in)
				//iin, _ := strconv.Atoi(in)
				v[string(ch)] = in
			}
		}
	}
	return v
}

func main() {
	expStr := getExpStr()
	v := getValue(expStr)
	cal := calculator.NewCalculator(expStr)
	fmt.Println("计算结果为: ", expStr, "=", cal.Run(v))
}
