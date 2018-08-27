package group

import "fmt"

type codeGroup struct {}

func NewCodeGroup() *codeGroup {
	return new(codeGroup)
}

func (r codeGroup)Find() {
	fmt.Println("find requirement group")
}

func (r codeGroup)Add() {
	fmt.Println("Add an requirement")
}

func (r codeGroup)Delete() {
	fmt.Println("delete an requirement")
}

func (r codeGroup)Change() {
	fmt.Println("change an requirement")
}

func (r codeGroup)Plan() {
	fmt.Println("change requirement group plan")
}
