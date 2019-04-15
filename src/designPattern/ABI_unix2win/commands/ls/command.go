package ls

import "designPattern/ABI_unix2win/commands"

type LSCommand struct {
	*commands.Command
}

func NewLSCommand() *LSCommand {
	return &LSCommand{new(commands.Command)}
}

func (c *LSCommand) Execute(vo commands.CommandVO) string {
	first := c.Command.BuildChain(NewLS(), NewLSA(), NewLSL())[0]
	return first.HandleMessage(vo)
}
