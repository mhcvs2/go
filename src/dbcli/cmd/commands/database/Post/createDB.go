package startscheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "createDB [-name=***]",
	Short:     "start scheduled backup",
	Long: `
start scheduled backup
`,
	Run:    createDB,
}
var (
	db = mgmt.NewDb()
	name string
)

func init() {
	CmdRun.Flag.StringVar(&name, "name", "", "Specify a database name.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func createDB(cmd *commands.Command, args []string) int {
	if name == "" {
		utils.PrintAndExit("lack name", utils.FailedTemplate)
	}
	resultBytes := db.Post(name)
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}