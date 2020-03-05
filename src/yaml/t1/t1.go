package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

type K8sObject struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   map[string]string `json:"metadata"`
}

func t1() {
	config, err := ioutil.ReadFile("D:\\github\\go\\src\\yaml\\t1\\cm.yaml")
	var k1 K8sObject
	err = yaml.Unmarshal(config, &k1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(k1)
}

func main() {
	t1()
}
