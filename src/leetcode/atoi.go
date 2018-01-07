package main

import "fmt"

func getStart(str string) (res int) {
	res = -1
	for i, c:= range str {
		res = i
		if !(c == ' ' || c == '0') {
			break
		}
	}
	return
}

func check_res(res int, fu bool) (int, bool) {
	if fu {
		if res > 2147483648 {
			res = 2147483648
			return res, true
		}
	} else {
		if res > 2147483647 {
			res = 2147483647
			return res, true
		}
	}
	return res, false
}

func myAtoi(str string) int {
	l := len(str)
	start := getStart(str)
	if start == -1 {
		return 0
	}
	fu := false
	if str[start] == '-'{
		fu = true
		start++
	} else if (str[start] == '+') {
		start++
	}
	cheng := 1
	res := 0
	for i := l-1; i>=start; i-- {
		switch str[i] {
		case '0':
			cheng *= 10
		case '1':
			res += 1 * cheng
			cheng *= 10
		case '2':
			res += 2 * cheng
			cheng *= 10
		case '3':
			res += 3 * cheng
			cheng *= 10
		case '4':
			res += 4 * cheng
			cheng *= 10
		case '5':
			res += 5 * cheng
			cheng *= 10
		case '6':
			res += 6 * cheng
			cheng *= 10
		case '7':
			res += 7 * cheng
			cheng *= 10
		case '8':
			res += 8 * cheng
			cheng *= 10
		case '9':
			res += 9 * cheng
			cheng *= 10
		default:
			cheng = 1
			res  = 0
		}
		if res2, ok := check_res(res, fu); ok {
			res = res2
			continue
		}
	}
	if res < 0 {
		res = res * -1
	}
	if fu {
		if res > 2147483648 {res = 2147483648}
		return res * -1
	}
	if res > 2147483647 {res = 2147483647}
	return res
}

func main() {
	fmt.Println(myAtoi("  +b12102370352"))
}