package GetAll
import (
	"dbcli/cmd/commands"
	"dbcli/mgmt"
	"encoding/json"
	myLogger "dbcli/logger"
	"dbcli/utils"
)

var CmdRun = &commands.Command{
	UsageLine: "listBackup",
	Short:     "list all backup",
	Long: `
list all backup
`,
	Run:    listBackup,
}
var (
	backup = mgmt.NewCluster()
)

type Result struct {
	Backup []string
}

var listbackupTemplate = `--------backups----------
{{ range $k,$v := . }}{{ $v }}
{{$v | fen}}
{{ end }}`

func init() {
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func listBackup(cmd *commands.Command, args []string) int {
	resultBytes := backup.GetAll()
	myLogger.Log.Debugf(string(resultBytes))
	var result  = Result{}
	err := json.Unmarshal(resultBytes, &result)
	if err != nil {
		myLogger.Log.Fatalf("fail to parse result:", err)
	}
	if len(result.Backup) == 0 {
		utils.PrintAndExit("no backups", utils.FailedTemplate)
	}
	utils.Tmpl(listbackupTemplate, result.Backup)
	return 0
}