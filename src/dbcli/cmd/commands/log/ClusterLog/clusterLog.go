package listbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "clusterLog",
	Short:     "get cluster logs",
	Long: `
get cluster logs
`,
	Run:    clusterLog,
}
var (
	clusterlog = mgmt.NewClusterLog()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func clusterLog(cmd *commands.Command, args []string) int {
	resultBytes := clusterlog.ClusterLog()
	utils.PrintAndExit(string(resultBytes), utils.SuccessTemplate)
	return 0
}