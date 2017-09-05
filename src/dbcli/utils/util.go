package utils

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"
	myLogger "dbcli/logger"
	"dbcli/logger/colors"
)

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
func Go(f func() error) chan error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return ch
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CloseFile attempts to close the passed file
// or panics with the actual error
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// __FILE__ returns the file name in which the function was invoked
func FILE() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// BeeFuncMap returns a FuncMap of functions used in different templates.
func BeeFuncMap() template.FuncMap {
	return template.FuncMap{
		"trim":       strings.TrimSpace,
		"bold":       colors.Bold,
		"headline":   colors.MagentaBold,
		"foldername": colors.RedBold,
		"endline":    EndLine,
		"tmpltostr":  TmplToString,
		"fen":        fenString,
		"red":        colors.Red,
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

func Tmpl(text string, data interface{}) {
	output := colors.NewColorWriter(os.Stderr)

	t := template.New("Usage").Funcs(BeeFuncMap())
	template.Must(t.Parse(text))

	err := t.Execute(output, data)
	if err != nil {
		myLogger.Log.Error(err.Error())
	}
}

func PrintErrorAndExit(message, errorTemplate string) {
	Tmpl(fmt.Sprintf(errorTemplate, message), nil)
	os.Exit(2)
}

func fenString(str string) string {
	r := "-"
	for i:=1; i<len([]rune(str));i++ {
		r += "-"
	}
	return r
}

func PrintAndExit(message, template string) {
	Tmpl(template, message)
	os.Exit(0)
}

var SuccessTemplate = `{{ . }}{{ endline }}`

var FailedTemplate = `{{ . | red }}{{ endline }}`

var preRunTemplate = ` {{ "Lack env 'MGMTSERVER' 'APPID' 'APPSCERET'" | headline }}
 For example:
		{{ "export MGMTSERVER=http://127.0.0.1:8080" | bold}}
		{{ "export APPID=test" | bold}}
		{{ "export APPSCERET=ee1d7336-73e3-452b-9e63-fdeaf2dccde6" | bold}}
`