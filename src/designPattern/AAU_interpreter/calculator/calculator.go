package calculator

import (
	cstack "github.com/mhcvs2/godatastructure/stack"
	"designPattern/AAU_interpreter/expressions"
)

type Calculator struct {
}

func NewCalculator(expStr string) *Calculator {
	stack := cstack.NewStack()
	var left expressions.Expression
	var right expressions.Expression
	for i:=0; i<len(expStr); i++ {
		switch expStr[i] {
		case '+':
			left = stack.Pop().(expressions.Expression)
			right = expressions.NewVarExpression(strconv.)
		}
	}
}
