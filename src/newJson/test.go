package main

import (
	"io/ioutil"
	"fmt"
	"github.com/json-iterator/go"
)

func tt1() {
	data, _ := ioutil.ReadFile("/root/github/go/src/newJson/test.json")
	fmt.Println(jsoniter.Get(data, "apiVersion").ToString())
	fmt.Println(jsoniter.Get(data, "items", 0, "apiVersion").ToString())
	fmt.Println(jsoniter.Get(data, "items", 0, "status","conditions", 0, "type").ToString())
	fmt.Println(jsoniter.Get(data, "items", 0, "spec","template", "spec", "containers", 0, "image").ToString())
	fmt.Println(jsoniter.Get(data, "items", 0, "status","startTime").ToString())
	fmt.Println(jsoniter.Get(data, "items", 0, "status","startTime").ToString())
	fmt.Println(jsoniter.Get(data, "items", 0, "status").ToString())
}

func tt2() {
	data, _ := ioutil.ReadFile("/root/github/go/src/newJson/test.json")
	state := "start"
	for i:=0; state!=""; i++ {
		state = jsoniter.Get(data, "items", i, "status","conditions", 0, "type").ToString()
		fmt.Println(state)
	}
	fmt.Println("done")

}

func main() {
	tt1()
}