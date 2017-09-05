package listbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "backup",
	Short:     "start backup",
	Long: `
start backup
`,
	Run:    backup,
}
var (
	cluster = mgmt.NewCluster()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func backup(cmd *commands.Command, args []string) int {
	resultBytes := cluster.StartBackup()
	utils.PrintAndExit(string(resultBytes), utils.SuccessTemplate)
	return 0
}