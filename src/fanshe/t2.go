package main

import "fmt"

type Mhc struct {
	A string
}

func Bianbie(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case Mhc:
		fmt.Println("struct")
	case *Mhc:
		fmt.Println("struct pointer")
	case map[string]string:
		fmt.Println("map[string]string")
	case map[string]interface{}:
		fmt.Println("map[string]interface")
	case []string:
		fmt.Println("[]string")
	case []interface{}:
		fmt.Println("[]interface")
	default:
		fmt.Println("not parse")
	}
}

func t21() {
	Bianbie(1)
	Bianbie(0.2)
	Bianbie("aaaa")
	Bianbie(Mhc{})
	Bianbie(&Mhc{})

	a := make(map[string]string)
	Bianbie(a)
	b := make(map[string]int)
	Bianbie(b)

	Bianbie([]string{"1", "2"})
	Bianbie([]int{1, 2})
}

//int
//float64
//string
//struct
//struct pointer
//map[string]string
//not parse
//[]string
//not parse

func main() {
	t21()
}
