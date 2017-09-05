package startscheduledbackup
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "startScheduledBackup [-interval=*m/h/d] [-time=**:**]",
	Short:     "start scheduled backup",
	Long: `
start scheduled backup
`,
	Run:    startScheduledBackup,
}
var (
	cluster = mgmt.NewCluster()
	interval string
	starttime string
)

func init() {
	CmdRun.Flag.StringVar(&interval, "interval", "1h", "Specify scheduled backup interval.")
	CmdRun.Flag.StringVar(&starttime, "time", "00:00", "Specify scheduled backup start time.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func startScheduledBackup(cmd *commands.Command, args []string) int {
	resultBytes := cluster.Start(interval, starttime)
	utils.Tmpl(utils.SuccessTemplate, string(resultBytes))
	return 0
}