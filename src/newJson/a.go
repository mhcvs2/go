package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

type ColorGroup struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Colors []string `json:"colors"`
}

func t1() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, _ := jsoniter.Marshal(group)
	fmt.Println(string(b))
}
//{"id":1,"name":"Reds","colors":["Crimson","Red","Ruby","Maroon"]}

func t2() {
	jsonText := `{"id":1,"name":"Reds","colors":["Crimson","Red","Ruby","Maroon"]}`
	r := ColorGroup{}
	jsoniter.Unmarshal([]byte(jsonText), &r)
	fmt.Println(r.ID)
	fmt.Println(r.Name)
	fmt.Println(r.Colors)
}
//1
//Reds
//[Crimson Red Ruby Maroon]

func t3() {
	jsonText := `{"id":1,"name":"Reds","colors":["Crimson","Red","Ruby","Maroon"]}`
	jsonByte := []byte(jsonText)
	fmt.Println(jsoniter.Get(jsonByte).ToString())
	fmt.Println(jsoniter.Get(jsonByte, "id").ToString())
	fmt.Println(jsoniter.Get(jsonByte, "id").ToInt())
	fmt.Println(jsoniter.Get(jsonByte, "name").ToString())
	fmt.Println(jsoniter.Get(jsonByte, "colors").ToString())
	fmt.Println(jsoniter.Get(jsonByte, "colors", 0).ToString())
	fmt.Println(jsoniter.Get(jsonByte, "colors", 4).ToString())
	fmt.Println(jsoniter.Get(jsonByte, "colors", 3).ToString())
	//{"id":1,"name":"Reds","colors":["Crimson","Red","Ruby","Maroon"]}
	//1
	//1
	//Reds
	//["Crimson","Red","Ruby","Maroon"]
	//Crimson
	//
	//Maroon
}

func t4() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := json.Marshal(group)
	fmt.Println(string(b))
}
//{"id":1,"name":"Reds","colors":["Crimson","Red","Ruby","Maroon"]}


func main() {
	t4()
}
