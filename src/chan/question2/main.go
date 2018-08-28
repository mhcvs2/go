package main

import (
	"bufio"
	"os"
	"fmt"
)

var prompt = "Enter a shi zi, e.g., 5 + 2, or ctrl+D to quit."

type Shizi struct {
	One int
	Two int
	Operation string
}

type Result struct {
	Data int
	ErrMsg string
}

func CreateSolver(question chan Shizi) chan Result {
	answer := make(chan Result)
	go func() {
		for{
			shizi := <- question
			res := Result{}
			switch shizi.Operation {
			case "+":
				res.Data = shizi.One + shizi.Two
			case "-":
				res.Data = shizi.One - shizi.Two
			case "*":
				res.Data = shizi.One * shizi.Two
			case "/":
				res.Data = shizi.One / shizi.Two
			default:
				res.ErrMsg = "invalid operation"
			}
			answer <- res
		}
	}()
	return answer
}

func interact(question chan Shizi, answer chan Result) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Println("shi zi:")
		line, _ := reader.ReadString('\n')
		var one, two int
		var operation string
		if _, err := fmt.Sscanf(line, "%d %s %d", &one, &operation, &two); err != nil {
			fmt.Fprintf(os.Stderr, "Invalid input\n")
			fmt.Println(err)
			continue
		}
		shizi := Shizi{
			One:one,
			Two:two,
			Operation:operation,
		}
		question <- shizi
		res := <- answer
		if res.ErrMsg != "" {
			fmt.Fprintf(os.Stderr, res.ErrMsg)
			fmt.Println()
		} else {
			fmt.Printf("result is: %d\n", res.Data)
		}
	}

}

func main() {
	question := make(chan Shizi)
	answer := CreateSolver(question)
	interact(question, answer)
}
