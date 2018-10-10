package main

import "fmt"

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = &Stduent{}   //必须加&, 否则Stduent does not implement People (Speak method has pointer receiver)
	think := "bitch"
	fmt.Println(peo.Speak(think))
}