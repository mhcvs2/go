package main

import (
	"designPattern/ABI_unix2win/commands"
	"designPattern/ABI_unix2win/commands/ls"
	"fmt"
)


var CommandEnumMap map[string]commands.ICommand

func init() {
	CommandEnumMap = make(map[string]commands.ICommand)
	CommandEnumMap["ls"] = ls.NewLSCommand()
}


func Exec(commandStr string) string {
	res := ""
	vo := commands.NewCommandVO(commandStr)
	if command, ok := CommandEnumMap[vo.CommandName()]; ok {
		res = command.Execute(*vo)
	} else {
		res = "failed to execute command"
	}
	return res
}

func main() {
	fmt.Println(Exec("ls -l"))
}