package commands

import "designPattern/command/group"

type ICommand interface {
	Execute()
}

type baseCommand struct {
	rg group.IGroup
	cg group.IGroup
}

func NewBaseCommand() *baseCommand {
	return &baseCommand{
		rg:group.NewRequirementGroup(),
		cg:group.NewCodeGroup(),
	}
}

func (c baseCommand)Execute() {

}

//-----------------------------------------

type command1 struct {
	baseCommand
}

func NewCommand1() *command1 {
	return &command1{*NewBaseCommand()}
}

func (c command1)Execute() {
	c.rg.Find()
	c.cg.Find()
	c.rg.Add()
	c.cg.Add()
	c.rg.Plan()
	c.cg.Plan()
}

//-------------------------------------------------


type command2 struct {
	baseCommand
}

func NewCommand2() *command2 {
	return &command2{*NewBaseCommand()}
}

func (c command2)Execute() {
	c.rg.Find()
	c.cg.Find()
	c.rg.Delete()
	c.cg.Delete()
	c.rg.Plan()
	c.cg.Plan()
}
