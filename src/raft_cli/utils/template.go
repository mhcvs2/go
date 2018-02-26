package utils

import (
	"bytes"
	myLogger "github.com/dbelt/dbcli/logger"
	"github.com/dbelt/dbcli/logger/colors"
	"os"
	"strings"
	"text/template"
)

// BeeFuncMap returns a FuncMap of functions used in different templates.
func BeeFuncMap() template.FuncMap {
	return template.FuncMap{
		"trim":       strings.TrimSpace,
		"bold":       colors.Bold,
		"headline":   colors.MagentaBold,
		"foldername": colors.RedBold,
		"endline":    EndLine,
		"tmpltostr":  TmplToString,
		"red":        colors.Red,
		"blue":       colors.Blue,
		"app":        GetAppName,
	}
}

// TmplToString parses a text template and return the result as a string.
func TmplToString(tmpl string, data interface{}) string {
	t := template.New("tmpl").Funcs(BeeFuncMap())
	template.Must(t.Parse(tmpl))

	var doc bytes.Buffer
	err := t.Execute(&doc, data)
	MustCheck(err)

	return doc.String()
}

// EndLine returns the a newline escape character
func EndLine() string {
	return "\n"
}

//Tmpl is template function
func Tmpl(text string, data interface{}) {
	output := colors.NewColorWriter(os.Stderr)

	t := template.New("Usage").Funcs(BeeFuncMap())
	template.Must(t.Parse(text))

	err := t.Execute(output, data)
	if err != nil {
		myLogger.Log.Error(err.Error())
	}
}

//PrintAndExit print error and exit program
func PrintAndExit(message, template string) {
	Tmpl(template, message)
	os.Exit(0)
}

//SuccessTemplate is success template
var SuccessTemplate = `{{ . | blue }}{{ endline }}`

//FailedTemplate is FailedTemplate
var FailedTemplate = `{{ . | red }}{{ endline }}`

//PreRunTemplate is PreRunTemplate
var PreRunTemplate = `{{ "Please specify 'APPID' and 'APPSCERET' before operation" | red }}
By:
a) Flag: --appid  --appsecret  --mgmt  --ops
b) Env:
      export APPID=test
      export APPSECRET=ee1d7336-73e3-452b-9e63-fdeaf2dccde6
      export MGMTSERVER=127.0.0.1:8080
      export OPSSERVER=127.0.0.1:8181
      export OUTPUT=json
c) A config file "{{ app }}.json/yml" in /etc or current directory
`
