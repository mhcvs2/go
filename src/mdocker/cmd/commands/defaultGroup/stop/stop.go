package stop
import (
	"mdocker/cmd/commands"
	"mdocker/models/v1"
)

var CmdRun = &commands.Command{
	UsageLine: "stop",
	Short:     "stop container",
	Long: `
stop container
`,
	Run:    stop,
}

var dm = docker.NewDockerManage()

func init() {
	commands.AddGroup("default", CmdRun)
}

func stop(cmd *commands.Command, args []string) int {
	dm.Stop(args, true)
	return 0
}