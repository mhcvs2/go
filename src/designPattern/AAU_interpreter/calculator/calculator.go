package calculator

import (
	"designPattern/AAU_interpreter/expressions"
	"github.com/godatastructure/stack"
)

type Calculator struct {
	expression expressions.Expression
}

func NewCalculator(expStr string) *Calculator {
	c := new(Calculator)
	s := stack.NewStack()
	charArray := []byte(expStr)
	var left expressions.Expression
	var right expressions.Expression
	for i:=0; i<len(charArray); i++{
		switch charArray[i] {
		case '+':
			left = s.Pop().(expressions.Expression)
			right = expressions.NewVarExpression(string(charArray[i+1]))
			i += 1
			s.Push(expressions.NewAddExpression(left, right))
		case '-':
			left = s.Pop().(expressions.Expression)
			right = expressions.NewVarExpression(string(charArray[i+1]))
			i += 1
			s.Push(expressions.NewSubExpression(left, right))
		default:
			s.Push(expressions.NewVarExpression(string(charArray[i])))
		}
	}
	c.expression = s.Pop().(expressions.Expression)
	return c
}

func (c *Calculator) Run(v map[string]int) int {
	return c.expression.Interpreter(v)
}
