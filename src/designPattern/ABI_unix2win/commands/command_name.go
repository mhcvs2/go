package commands

import "github.com/bitrise-io/go-utils/sliceutil"

type ICommandName interface {
	HandleMessage(vo CommandVO) string
	SetNext(operator ICommandName)
	OperateParam() string
	Echo(vo CommandVO) string
}

type CommandName struct {
	nextOperator ICommandName
	processor ICommandName
}

func NewCommandName(processor ICommandName) *CommandName {
	return &CommandName{processor: processor}
}

func (c *CommandName) HandleMessage(vo CommandVO) string {
	res := ""
	if len(vo.paramList) == 0 || sliceutil.IsStringInSlice(c.processor.OperateParam(), vo.paramList) {
		res = c.processor.Echo(vo)
	} else {
		if c.nextOperator != nil {
			res =     c.nextOperator.HandleMessage(vo)
		} else {
			res = "failed to execute command"
		}
	}
	return res
}

func (c *CommandName) SetNext(operator ICommandName) {
	c.nextOperator = operator
}

func (c *CommandName) OperateParam() string {
	panic("implement me")
}

func (c *CommandName) Echo(vo CommandVO) string {
	panic("implement me")
}




