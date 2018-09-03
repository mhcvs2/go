package strategy

type Context struct {
	strategy IStrategy
}

func NewContext(strategy IStrategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context)Operate() {
	c.strategy.operate()
}

