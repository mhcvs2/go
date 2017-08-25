package getip

import (
	"dockerManage/cmd/commands"
	myLogger "dockerManage/logger"
	"fmt"
	"strings"
	"dockerManage/utils"
	"os"
	"strconv"
)

var CmdRun = &commands.Command{
	UsageLine: "getip name1 name2...\n",
	Short:     "get ip by container id",
	Long: `
get ip by container id
`,
	Run:    getIp,
}

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func getIp(cmd *commands.Command, args []string) int {
	//started <- true

	if len(args) == 0 {
		ids, names := idsNames()
		ips := getips(ids)
		out(ips, ids, names)
	} else {
		ips := getips(args)
		out(ips, args)

	}

	return 0

}

func getips(ids []string) (ips []string) {
	for _, name := range ids {
		out, err := utils.Exec_shell("docker inspect --format '{{ .NetworkSettings.IPAddress }}' " + name)
		if err != nil {
			myLogger.Log.Error("Get ip of " + name + "fail: " + err.Error())
			continue
		}
		out = strings.TrimSpace(out)
		ips = append(ips, out)
	}
	return ips
}

func idsNames() (ids, names []string) {
	out, err := utils.Exec_shell("docker ps|wc -l")
	if err != nil {
		myLogger.Log.Error(err.Error())
		return ids, names
	}
	if outInt, _ := strconv.Atoi(strings.TrimSpace(out)); outInt == 0 {
		fmt.Println("no alive docker container")
		return ids, names
	}

	out, err = utils.Exec_shell("docker ps | awk '{print $1}'")
	if err != nil {
		myLogger.Log.Error("docker ps err: " + err.Error())
		os.Exit(1)
	}
	out = strings.TrimSpace(out)
	ids = strings.Split(out, "\n")[1:]

	out, err = utils.Exec_shell("docker ps | awk '{print $NF}'")
	if err != nil {
		myLogger.Log.Error("docker ps err: " + err.Error())
		os.Exit(1)
	}
	out = strings.TrimSpace(out)
	names = strings.Split(out, "\n")[1:]
	return ids, names
}

func out(ips [] string, namess ...[]string) {
	printMessage := ""
	for i :=0; i < len(ips); i++ {
		printMessage += ips[i]+"  "
		if len(namess) > 0 {
			for j := 0; j < len(namess); j++ {
				printMessage += namess[j][i] +"  "
			}
			printMessage += "\n"
		}
	}
	fmt.Print(printMessage)
}