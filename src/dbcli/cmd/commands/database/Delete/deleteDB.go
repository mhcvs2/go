package startscheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "deleteDB [-name=***]",
	Short:     "delete a database from master",
	Long: `
delete a backup
`,
	Run:    deleteDB,
}
var (
	db = mgmt.NewDb()
	name string
)

func init() {
	CmdRun.Flag.StringVar(&name, "name", "", "Specify a database name to delete.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func deleteDB(cmd *commands.Command, args []string) int {
	if name == "" {
		utils.PrintAndExit("lack name", utils.FailedTemplate)
	}
	resultBytes := db.Delete(name)
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}