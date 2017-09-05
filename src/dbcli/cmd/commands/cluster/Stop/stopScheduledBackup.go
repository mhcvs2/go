package scheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "stopScheduledBackup",
	Short:     "stop scheduled backup",
	Long: `
stop scheduled backup
`,
	Run:    stopScheduledBackup,
}
var (
	cluster = mgmt.NewCluster()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func stopScheduledBackup(cmd *commands.Command, args []string) int {
	resultBytes := cluster.Stop()
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}