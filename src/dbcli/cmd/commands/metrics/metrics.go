package listbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "getmetrics",
	Short:     "get metrics",
	Long: `
get metrics
`,
	Run:    getmetrics,
}
var (
	metrics = mgmt.NewMetrics()
)

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func getmetrics(cmd *commands.Command, args []string) int {
	resultBytes := metrics.Get()
	utils.PrintAndExit(string(resultBytes), utils.SuccessTemplate)
	return 0
}