package listbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "startCluster",
	Short:     "start cluster",
	Long: `
start cluster
`,
	Run:    startCluster,
}
var (
	cluster = mgmt.NewCluster()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func startCluster(cmd *commands.Command, args []string) int {
	resultBytes := cluster.StartCluster()
	utils.PrintAndExit(string(resultBytes), utils.SuccessTemplate)
	return 0
}