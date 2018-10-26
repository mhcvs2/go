package expressions

type Expression interface {
	Interpreter(vars map[string]int) int
}

////////////////////////////////////////////////////////////

type VarExpression struct {
	Key string
}

func NewVarExpression(key string) *VarExpression {
	return &VarExpression{Key: key}
}

func (v VarExpression) Interpreter(vars map[string]int) int {
	if res, ok := vars[v.Key]; ok{
		return res
	} else {
		panic("error, value is nil")
	}
}

////////////////////////////////////////////////////////////

type SymbolExpression struct {
	Left Expression
	Right Expression
}

func NewSymbolExpression(left, right Expression) *SymbolExpression {
	return &SymbolExpression{
		Left: left,
		Right: right,
	}
}

func (v SymbolExpression) Interpreter(vars map[string]int) int {
	panic("not implement")
}

///////////////////////////////////////////////////////////////////////////

type AddExpression struct {
	SymbolExpression
}

func NewAddExpression(left, right Expression) *AddExpression {
	return &AddExpression{
		*NewSymbolExpression(left, right),
	}
}

func (v AddExpression) Interpreter(vars map[string]int) int {
	return v.Left.Interpreter(vars) + v.Right.Interpreter(vars)
}

///////////////////////////////////////////////////////////////////////////

type SubExpression struct {
	SymbolExpression
}

func NewSubExpression(left, right Expression) *SubExpression {
	return &SubExpression{
		*NewSymbolExpression(left, right),
	}
}

func (v SubExpression) Interpreter(vars map[string]int) int {
	return v.Left.Interpreter(vars) - v.Right.Interpreter(vars)
}
