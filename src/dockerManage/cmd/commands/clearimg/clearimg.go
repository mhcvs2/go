package clearimg

import (
	"dockerManage/cmd/commands"
	myLogger "dockerManage/logger"
	"dockerManage/utils"
	"strconv"
	"fmt"
)

var CmdRun = &commands.Command{
	UsageLine: "clearimg",
	Short:     "clear temporary images",
	Long: `
clear temporary images
`,
	Run:    clearImg,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func clearImg(cmd *commands.Command, args []string) int {
	//started <- true
	out, err := utils.Exec_shell("docker images -q -f dangling=true|wc -l")
	if err != nil {
		myLogger.Log.Error(err.Error())
		return 0
	}
	if outInt, _ := strconv.Atoi(out); outInt == 0 {
		fmt.Println("no images to clear")
		return 0
	}

	out, err = utils.Exec_shell("docker rmi $(docker images -q -f dangling=true)")
	if err != nil {
		myLogger.Log.Error(err.Error())
		return 0
	}
	myLogger.Log.Info(out)

	return 0
}