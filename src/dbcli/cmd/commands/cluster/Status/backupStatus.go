package listbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "backupStatus",
	Short:     "get backup status",
	Long: `
get backup status
`,
	Run:    backupStatus,
}
var (
	cluster = mgmt.NewCluster()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func backupStatus(cmd *commands.Command, args []string) int {
	resultBytes := cluster.Status()
	utils.PrintAndExit(string(resultBytes), utils.SuccessTemplate)
	return 0
}