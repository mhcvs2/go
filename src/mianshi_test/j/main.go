package main

import (
	"fmt"
	"reflect"
)

func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn3 := struct {
		name string
		age  int
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	//错
	//if sn1 == sn3 {
	//	fmt.Println("sn1 == sn3")
	//}
	if reflect.DeepEqual(sn1, sn3) {
		fmt.Println("sn1 == sn3")
	} else {
		fmt.Println("sn1 != sn3")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	// 错的
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}
}
//sn1 == sn2
//sn1 != sn3
//sm1 == sm2