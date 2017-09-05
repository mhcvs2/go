package startscheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "startRestore [-name=***]",
	Short:     "start restore",
	Long: `
start restore
`,
	Run:    startRestore,
}
var (
	cluster = mgmt.NewCluster()
	name string
)

func init() {
	CmdRun.Flag.StringVar(&name, "name", "", "Specify a backup file name to restore.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func startRestore(cmd *commands.Command, args []string) int {
	if name == "" {
		utils.PrintAndExit("lack name", utils.FailedTemplate)
	}
	resultBytes := cluster.StartRestore(name)
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}