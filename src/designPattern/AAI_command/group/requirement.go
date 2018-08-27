package group

import "fmt"

type requirementGroup struct {}

func NewRequirementGroup() *requirementGroup {
	return new(requirementGroup)
}

func (r requirementGroup)Find() {
	fmt.Println("find requirement group")
}

func (r requirementGroup)Add() {
	fmt.Println("Add an requirement")
}

func (r requirementGroup)Delete() {
	fmt.Println("delete an requirement")
}

func (r requirementGroup)Change() {
	fmt.Println("change an requirement")
}

func (r requirementGroup)Plan() {
	fmt.Println("change requirement group plan")
}