package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)

	fmt.Println()

	for n, animal := range animals {
		fmt.Println("第 " + strconv.Itoa(n) + " 个")
		fmt.Println("Order: " + animal.Order + ", Name: " + animal.Name)
	}

}

//[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
//第 0 个
//Order: Monotremata, Name: Platypus
//第 1 个
//Order: Dasyuromorphia, Name: Quoll