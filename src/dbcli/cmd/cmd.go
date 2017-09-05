package cmd

import (
	"dbcli/cmd/commands"
	"dbcli/utils"
	_ "dbcli/cmd/commands/cluster/Delete"
	_ "dbcli/cmd/commands/cluster/DeleteAll"
	_ "dbcli/cmd/commands/cluster/GetAll"
	_ "dbcli/cmd/commands/cluster/ShowCluster"
	_ "dbcli/cmd/commands/cluster/Start"
	_ "dbcli/cmd/commands/cluster/StartBackup"
	_ "dbcli/cmd/commands/cluster/StartCluster"
	_ "dbcli/cmd/commands/cluster/StartRestore"
	_ "dbcli/cmd/commands/cluster/Status"
	_ "dbcli/cmd/commands/cluster/Stop"
	_ "dbcli/cmd/commands/cluster/StopCluster"
	_ "dbcli/cmd/commands/metrics"
	_ "dbcli/cmd/commands/database/Post"
	_ "dbcli/cmd/commands/database/Delete"
	_ "dbcli/cmd/commands/dbuser/Post"
	_ "dbcli/cmd/commands/dbuser/Delete"
	_ "dbcli/cmd/commands/log/ClusterLog"
	_ "dbcli/cmd/commands/log/GetDockerContainerLog"
)

var usageTemplate = `dbcli is a Fast and Flexible tool for managing your DB cluster.

{{"USAGE" | headline}}
    {{"dbcli command [arguments]" | bold}}

{{"AVAILABLE COMMANDS" | headline}}
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-20s" | bold}} {{.Short}}{{end}}{{end}}

Use {{"dbcli help [command]" | bold}} for more information about a command.

{{"ADDITIONAL HELP TOPICS" | headline}}
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-20s"}} {{.Short}}{{end}}{{end}}

Use {{"dbcli help [topic]" | bold}} for more information about that topic.
`

var helpTemplate = `{{"USAGE" | headline}}
  {{.UsageLine | printf "dbcli %s" | bold}}
{{if .Options}}{{endline}}{{"OPTIONS" | headline}}{{range $k,$v := .Options}}
  {{$k | printf "-%s" | bold}}
      {{$v}}
  {{end}}{{end}}
{{"DESCRIPTION" | headline}}
  {{tmpltostr .Long . | trim}}
`

var ErrorTemplate = `mdocker: %s.
Use {{"dbcli help" | bold}} for more information.
`

func Usage() {
	utils.Tmpl(usageTemplate, commands.AvailableCommands)
}

func Help(args []string) {
	if len(args) == 0 {
		Usage()
	}
	if len(args) != 1 {
		utils.PrintErrorAndExit("Too many arguments", ErrorTemplate)
	}

	arg := args[0]

	for _, cmd := range commands.AvailableCommands {
		if cmd.Name() == arg {
			utils.Tmpl(helpTemplate, cmd)
			return
		}
	}
	utils.PrintErrorAndExit("Unknown help topic", ErrorTemplate)
}
