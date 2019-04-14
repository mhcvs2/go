package commands

type ICommand interface {
	Execute(vo CommandVO) string
	BuildChain(commandNames ...ICommandName) []ICommandName
}

type Command struct {}

func (Command) Execute(vo CommandVO) string {
	panic("implement me")
}

func (Command) BuildChain(commandNames ...ICommandName) []ICommandName {
	res := make([]ICommandName, len(commandNames))
	for i:=0; i<len(commandNames); i++ {
		res[i] = commandNames[i]
		if i > 0 {
			res[i-1].SetNext(res[i])
		}
	}
	return res
}


