package listbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "getClusterInfo",
	Short:     "get cluster info",
	Long: `
get cluster info
`,
	Run:    getClusterInfo,
}
var (
	cluster = mgmt.NewCluster()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func getClusterInfo(cmd *commands.Command, args []string) int {
	resultBytes := cluster.ShowCluster()
	utils.PrintAndExit(string(resultBytes), utils.SuccessTemplate)
	return 0
}