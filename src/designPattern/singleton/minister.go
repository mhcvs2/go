package main

import "designPattern/singleton/emperor"

func main() {
	er := emperor.GetInstance()
	er.Name = "haha"
	er.Say()
	for i:=0; i<3; i++{
		er := emperor.GetInstance()
		er.Say()
	}
}
