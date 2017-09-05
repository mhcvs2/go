package listbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "nodeLog [-containerId=***]",
	Short:     "get a node log by containerId",
	Long: `
get cluster logs
`,
	Run:    nodeLog,
}
var (
	clusterlog = mgmt.NewClusterLog()
	containerId string
)

func init() {
	CmdRun.Flag.StringVar(&containerId, "containerId", "", "Specify a containerId.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func nodeLog(cmd *commands.Command, args []string) int {
	if containerId == ""{
		utils.PrintAndExit("lack containerId", utils.FailedTemplate)
	}
	resultBytes := clusterlog.GetDockerContainerLog(containerId)
	utils.PrintAndExit(string(resultBytes), utils.SuccessTemplate)
	return 0
}