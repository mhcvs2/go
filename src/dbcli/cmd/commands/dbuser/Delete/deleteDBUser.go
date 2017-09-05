package startscheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "deleteDBUser [-dbName=***] [-userName=***]",
	Short:     "delete a database and it's bind user",
	Long: `
delete a database and it's bind user
`,
	Run:    deleteDBUser,
}
var (
	dbu = mgmt.NewDbu()
	dbName string
	userName string
)

func init() {
	CmdRun.Flag.StringVar(&dbName, "dbName", "", "Specify a database name.")
	CmdRun.Flag.StringVar(&userName, "userName", "", "Specify a user name.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func deleteDBUser(cmd *commands.Command, args []string) int {
	if dbName == "" ||  userName == ""{
		utils.PrintAndExit("lack dbName/userName", utils.FailedTemplate)
	}
	resultBytes := dbu.Delete2(dbName, userName)
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}