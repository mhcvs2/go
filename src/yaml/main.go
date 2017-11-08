package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`  // Affects YAML field names too.
	Age int `json:"age"`
}

func t1() {
	//config, err := yaml.ReadFile("/root/github/go/src/yaml/conf.yaml")
	p := Person{"John", 30}
	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
}

func t2() {
	config, err := ioutil.ReadFile("/root/github/go/src/yaml/conf.yaml")
	var p2 Person
	err = yaml.Unmarshal(config, &p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(p2)
}

/*
- name: mhc
  age: 100

- name: mhc2
  age: 100

- name: mhc3
  age: 100
*/
func t3() {
	config, err := ioutil.ReadFile("/root/github/go/src/yaml/conf.yaml")
	var p2 []Person
	err = yaml.Unmarshal(config, &p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(p2)
}
//[{mhc 100} {mhc2 100} {mhc3 100}]

func main() {
	t3()
}