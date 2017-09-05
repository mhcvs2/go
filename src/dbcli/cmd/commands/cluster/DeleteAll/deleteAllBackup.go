package startscheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "deleteAllBackup [-name=***]",
	Short:     "delete all backup",
	Long: `
delete all backup
`,
	Run:    deleteAllBackup,
}
var (
	cluster = mgmt.NewCluster()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func deleteAllBackup(cmd *commands.Command, args []string) int {
	resultBytes := cluster.DeleteAll()
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}