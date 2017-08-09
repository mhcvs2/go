package main

import (
	"fmt"
	"runtime"
	"math"
	"bufio"
	"os"
)

type polar struct {
	radius float64
	seite float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle (in degree), e.g., 12.5 90, or %s to quit."

func init()  {
	if runtime.GOOS == "windows"{
		prompt = fmt.Sprintf(prompt, "Ctrl+z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			seite := polarCoord.seite
			x := polarCoord.radius * math.Cos(seite)
			y := polarCoord.radius * math.Sin(seite)
			answers <- cartesian{x,y}
		}
	}()
	return answers
}

const result = "Polar radius=%.02f seite=%.02f度 -> Cartesian x=%.02f y=%.02f\n"

func interact( questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for{
		fmt.Println("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil{
			break
		}
		var radius, seite float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &seite); err != nil{
			fmt.Println(os.Stderr, "invalid input")
			fmt.Println(err)
			continue
		}
		questions <- polar{radius, seite}
		coord := <- answers
		fmt.Printf(result, radius, seite, coord.x, coord.y)
	}
	fmt.Println()
}