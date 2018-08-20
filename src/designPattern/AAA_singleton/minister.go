package main

import (
	"designPattern/AAA_singleton/emperor2"
)

func main() {
	er := emperor2.GetInstance()
	er.Name = "haha"
	er.Say()
	for i:=0; i<3; i++{
		er := emperor2.GetInstance()
		er.Say()
	}
}
