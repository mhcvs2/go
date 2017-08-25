package main

import (
	myLogger "github.com/astaxie/beego/logs"
)

/*
const (
	AdapterConsole   = "console"
	AdapterFile      = "file"
	AdapterMultiFile = "multifile"
	AdapterMail      = "smtp"
	AdapterConn      = "conn"
	AdapterEs        = "es"
	AdapterJianLiao  = "jianliao"
	AdapterSlack     = "slack"
	AdapterAliLS     = "alils"
)

*/

func consoleTest() {
	log := myLogger.NewLogger(10000)
	log.SetLogger("console")
	log.Info("sdffsddsfdf")
	log.Error("sdfsdfsdfsssssssssssssss")
}

func fileTest() {
	log := myLogger.NewLogger(10000)
	//log.SetLogger("/root/github/go/src/logTest/test.log")
	//	{
	//	"filename":"logs/beego.log",
	//	"maxLines":10000,
	//	"maxsize":1024,
	//	"daily":true,
	//	"maxDays":15,
	//	"rotate":true,
	//  	"perm":"0600"
	//	}
	log.SetLogger("file","{\"filename\":\"/root/github/go/src/logTest/test.log\"}")
	log.Info("sdffsddsfdf")
	log.Error("sdfsdfsdfsssssssssssssss")
}

func main() {
	consoleTest()
}