package version
import (
	"gocli/cmd/commands"
	"gocli/utils"
	"fmt"
	"gocli/config"
)

var CmdRun = &commands.Command{
	UsageLine: "version/v",
	Short:     "show version",
	Long: `
show version
`,
	Run:    showVersion,
}

func init() {
	commands.AddGroup("default", CmdRun)
}

func showVersion(cmd *commands.Command, args []string) int {
	show := fmt.Sprintf("%s version %s", config.Conf.AppName, config.Conf.Version)
	utils.Tmpl(utils.SuccessTemplate, show)
	return 0
}