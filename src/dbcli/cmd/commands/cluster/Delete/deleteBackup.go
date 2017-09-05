package startscheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "deleteBackup [-name=***]",
	Short:     "delete a backup",
	Long: `
delete a backup
`,
	Run:    deleteBackup,
}
var (
	cluster = mgmt.NewCluster()
	name string
)

func init() {
	CmdRun.Flag.StringVar(&name, "name", "", "Specify a backup file name to delete.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func deleteBackup(cmd *commands.Command, args []string) int {
	if name == "" {
		utils.PrintAndExit("lack name", utils.FailedTemplate)
	}
	resultBytes := cluster.Delete(name)
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}