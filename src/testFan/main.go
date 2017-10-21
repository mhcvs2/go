package main

import (
	"fmt"
)

type Result struct {
	Code int
	Data interface{}
}

func main() {
	questions := make(chan Result)
	defer close(questions)
	createSolver(questions)
	questions <- Result{0, []string{"ss","sdfds"}}
}

func createSolver(questions chan Result){
	go func() {
		polarCoord := <-questions
		fmt.Println(polarCoord)
	}()
}

func interact( questions chan Result, answers chan Result) {
		questions <- Result{0, []string{"ss","sdfds"}}
		coord := <- answers
		fmt.Println(coord)
}
