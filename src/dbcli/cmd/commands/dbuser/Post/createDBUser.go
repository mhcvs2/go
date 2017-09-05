package startscheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "createDBUser [-dbName=***] [-userName=***] [-password=***]",
	Short:     "create a database and a bind user for it",
	Long: `
create a database and a bind user for it
`,
	Run:    createDBUser,
}
var (
	dbu = mgmt.NewDbu()
	dbName string
	userName string
	password string
)

func init() {
	CmdRun.Flag.StringVar(&dbName, "dbName", "", "Specify a database name.")
	CmdRun.Flag.StringVar(&userName, "userName", "", "Specify a user name.")
	CmdRun.Flag.StringVar(&password, "password", "", "Specify password for user.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func createDBUser(cmd *commands.Command, args []string) int {
	if dbName == "" ||  userName == "" || password==""{
		utils.PrintAndExit("lack dbName/userName/password", utils.FailedTemplate)
	}
	resultBytes := dbu.Post2(dbName, userName, password)
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}