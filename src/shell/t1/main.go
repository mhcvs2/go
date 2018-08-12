package main

import (
	"os/exec"
	"os"
	"fmt"
	log "github.com/sirupsen/logrus"
	"bytes"
)

func exec_shell1(s string) (error) {
	cmd := exec.Command("/bin/bash", "-c", s)

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func t1() {
	exec_shell1("ls /")
}

func t2() {
	os.Stdout.Write([]byte("hello world"))
}

type myWriter struct {}

func (w myWriter)Write(p []byte) (n int, err error){
	fmt.Println("==============================")
	fmt.Println(string(p))
	return 0, nil
}

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

func exec_shell3(s string) (error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	//out := &myWriter{}
	log.SetFormatter(&MyJSONFormatter{})
	cmdLog := *log.WithFields(log.Fields{"name":"aa"})
	cmd.Stdout = cmdLog.WriterLevel(log.InfoLevel)
	cmd.Stderr = cmdLog.WriterLevel(log.ErrorLevel)
	err := cmd.Start()
	if err != nil {
		return err
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return  nil
}

func t3() {
	exec_shell3("ls /")
}

func exec_shell4(s string) (error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	//out := &myWriter{}
	log.SetFormatter(&log.TextFormatter{ForceColors:true})
	cmdLog := *log.WithFields(log.Fields{"name":"aa"})
	cmd.Stdout = cmdLog.WriterLevel(log.InfoLevel)
	cmd.Stderr = cmdLog.WriterLevel(log.ErrorLevel)
	err := cmd.Start()
	if err != nil {
		return err
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return  nil
}

func t4() {
	exec_shell4("ls /")
}



func main() {
	t3()
}