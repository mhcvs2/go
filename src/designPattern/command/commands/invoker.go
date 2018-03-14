package commands

import "fmt"

type IInvoker interface {
	SetCommand(command ICommand)
	Action()
}

type Invoker struct {
	command ICommand
}

func NewInvoker() *Invoker {
	return new(Invoker)
}

func (ik *Invoker) SetCommand(command ICommand) {
	ik.command = command
}

func (ik *Invoker) Action() {
	fmt.Println("start action----------------------------")
	ik.command.Execute()
}
