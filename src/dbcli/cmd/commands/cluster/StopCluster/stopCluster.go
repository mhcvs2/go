package listbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "stopCluster",
	Short:     "stop cluster",
	Long: `
stop cluster
`,
	Run:    stopCluster,
}
var (
	cluster = mgmt.NewCluster()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func stopCluster(cmd *commands.Command, args []string) int {
	resultBytes := cluster.StopCluster()
	utils.PrintAndExit(string(resultBytes), utils.SuccessTemplate)
	return 0
}