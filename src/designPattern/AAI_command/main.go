package main

import "designPattern/AAI_command/commands"

func main() {

	xiao3 := commands.NewInvoker()
	xiao3.SetCommand(commands.NewCommand1())
	xiao3.Action()

	xiao3.SetCommand(commands.NewCommand2())
	xiao3.Action()
}

//start action----------------------------
//find requirement group
//find requirement group
//Add an requirement
//Add an requirement
//change requirement group plan
//change requirement group plan
//start action----------------------------
//find requirement group
//find requirement group
//delete an requirement
//delete an requirement
//change requirement group plan
//change requirement group plan
