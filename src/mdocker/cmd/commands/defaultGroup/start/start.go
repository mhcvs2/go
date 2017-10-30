package start
import (
	"mdocker/cmd/commands"
	"mdocker/models/v1"
)

var CmdRun = &commands.Command{
	UsageLine: "start",
	Short:     "start container",
	Long: `
start container
`,
	Run:    start,
}

var dm = docker.NewDockerManage()

func init() {
	commands.AddGroup("default", CmdRun)
}

func start(cmd *commands.Command, args []string) int {
	dm.Start(args, true)
	return 0
}