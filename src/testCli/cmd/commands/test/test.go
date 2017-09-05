package test

import (
	"testCli/cmd/commands"
	"testCli/logger/colors"
	"text/template"
	"os"
	myLogger "testCli/logger"
	"strings"
)

var CmdRun = &commands.Command{
	UsageLine: "test [name] [-main=...]",
	Short:     "test short",
	Long: `
test long........................

`,
	Run:    testRun,
}

var (
	main string
)

func lenOfString(str string) int  {
	return len([]rune(str))
}

func fenString(str string) string {
	r := "-"
	for i:=1; i<len([]rune(str));i++ {
		r += "-"
	}
	return r
}

func MyFuncMap() template.FuncMap {
	return template.FuncMap{
		"trim":       strings.TrimSpace,
		"bold":       colors.Bold,
		"headline":   colors.MagentaBold,
		"foldername": colors.RedBold,
		"len": lenOfString,
		"fen": fenString,
	}
}

func Tmpl(text string, data interface{}) {
	output := colors.NewColorWriter(os.Stderr)

	t := template.New("Usage").Funcs(MyFuncMap())
	template.Must(t.Parse(text))

	err := t.Execute(output, data)
	if err != nil {
		myLogger.Log.Error(err.Error())
	}
}

type Mhc struct {
	A string
	B string
}
var mhcs = []*Mhc{}

var testTemplate = `---{{ . | len}}---`
var testTemplate2 = `{{ range . }}{{ .A | printf "%-11s"}}: {{.B}}
{{ end }}`
var testTemplate3 = `---backups---
{{ range . }}{{ . }}
{{. | fen}}
{{ end }}`
var testTemplate4 = `---backups---
{{ range $k,$v := . }}{{$k}}.{{ $v }}
{{$v | fen}}
{{ end }}`

func init() {
	CmdRun.Flag.StringVar(&main, "main", "default_main","Specify main go files.")
	commands.AvailableCommands = append(commands.AvailableCommands, CmdRun)
}

func testRun(cmd *commands.Command, args []string) int {

	//mhc := Mhc{"aa","ss"}
	//mhc2 := Mhc{"aa222","ss222"}
	//mhcs = append(mhcs, &mhc)
	//mhcs = append(mhcs, &mhc2)
	//Tmpl(testTemplate2, mhcs)
	backs := []string{"hello-world", "hello-world","hello-world","hello-world","hello-world"}
	Tmpl(testTemplate4, backs)

	return 0

}