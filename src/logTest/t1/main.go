package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 36
	gray    = 37
)


type MyJSONFormatter struct {
}

func (f *MyJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	var levelColor int
	switch entry.Level {
	case log.ErrorLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	fmt.Fprintf(b, "\x1b[%dm%s:\x1b[%dm %s\n", gray, entry.Data["name"], levelColor, entry.Message)
	return b.Bytes(), nil
}

func t1() {
	log.SetFormatter(&MyJSONFormatter{})
	log.SetLevel(log.InfoLevel)
	cmdLog := *log.WithFields(log.Fields{"name":"aa"})
	cmdLog.Info("hahaha")
	cmdLog.Error("lalalal")
}
//aa: hahaha
//aa: lalalal


func t2() {
	log.SetFormatter(&log.TextFormatter{ForceColors:true})
	cmdLog := *log.WithFields(log.Fields{"name":"aa"})
	cmdLog.Info("hahaha")
	cmdLog.Error("lalalal")
}

func t3() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	cmdLog := *log.WithFields(log.Fields{"name":"aa"})
	cmdLog.Info("hahaha")
	cmdLog.Error("lalalal")
}
//time="2019-01-08T14:39:18+08:00" level=info msg=hahaha func=main.t3 file="D:/github/go/src/logTest/t1/main.go:60" name=aa
//time="2019-01-08T14:39:18+08:00" level=error msg=lalalal func=main.t3 file="D:/github/go/src/logTest/t1/main.go:61" name=aa

func main() {
	t3()
}