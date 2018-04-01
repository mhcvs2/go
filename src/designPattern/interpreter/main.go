package main

import (
	"github.com/mhcvs2/godatastructure/stack"
	"fmt"
	"strconv"
)

type Expression interface {
	Interpreter(v map[string]int) int
}
////////////////////////////////////////

type VarExpression struct {
	key string
}

func NewVarExpression(key string) *VarExpression {
	return &VarExpression{key:key}
}

func (e *VarExpression)Interpreter(v map[string]int) int {
	if i, ok := v[e.key]; ok {
		return i
	} else {
		return 0
	}
}

////////////////////////////////////////////////

type SymbolExpression struct {
	left Expression
	right Expression
}

func NewSymbolExpression(left, right Expression) *SymbolExpression {
	return &SymbolExpression{left:left, right:right}
}

//////////////////////////////////////////////////

type AddExpression struct {
	SymbolExpression
}

func NewAddExpression(left, right Expression) *AddExpression {
	return &AddExpression{
		SymbolExpression: *NewSymbolExpression(left, right),
	}
}

func (e *AddExpression) Interpreter(v map[string]int) int {
	return e.left.Interpreter(v) + e.right.Interpreter(v)
}


//////////////////////////////////////////////////

type SubExpression struct {
	SymbolExpression
}

func NewSubExpression(left, right Expression) *SubExpression {
	return &SubExpression{
		SymbolExpression: *NewSymbolExpression(left, right),
	}
}

func (e *SubExpression) Interpreter(v map[string]int) int {
	return e.left.Interpreter(v) - e.right.Interpreter(v)
}

////////////////////////////////////////////////////////

type Calculator struct {
	expression Expression
}

func NewCalculator(expStr string) *Calculator {
	c := new(Calculator)
	s := stack.NewStack()
	charArray := []byte(expStr)
	var left Expression
	var right Expression
	for i:=0; i<len(charArray); i++{
		switch charArray[i] {
		case '+':
			left = s.Pop().(Expression)
			right = NewVarExpression(string(charArray[i+1]))
			i += 1
			s.Push(NewAddExpression(left, right))
		case '-':
			left = s.Pop().(Expression)
			right = NewVarExpression(string(charArray[i+1]))
			i += 1
			s.Push(NewSubExpression(left, right))
		default:
			s.Push(NewVarExpression(string(charArray[i])))
		}
	}
	c.expression = s.Pop().(Expression)
	return c
}

func (c *Calculator) Run(v map[string]int) int {
	return c.expression.Interpreter(v)
}

//////////////////////////////////////////////////////////////

func getExpStr() string {
	fmt.Print("请输入表达式:")
	var expStr string
	fmt.Scanln(&expStr)
	return expStr
}

func getValue(expStr string)  map[string]int {
	v := make(map[string]int)
	for _, ch := range []byte(expStr){
		if ch != '+' && ch != '-'{
			if _, ok := v[string(ch)]; !ok {
				var in string
				fmt.Printf("输入%s的值: ", string(ch))
				fmt.Scanln(&in)
				iin, _ := strconv.Atoi(in)
				v[string(ch)] = iin
			}
		}
	}
	return v
}

func main() {
	expStr := getExpStr()
	v := getValue(expStr)
	cal := NewCalculator(expStr)
	fmt.Println("计算结果为: ", expStr, "=", cal.Run(v))
}
//
//请输入表达式:a+b-c
//输入a的值: 100
//输入b的值: 20
//输入c的值: 40
//计算结果为:  a+b-c = 80