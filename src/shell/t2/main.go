package main

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"fmt"
	"os/exec"
	"sync"
	"github.com/panjf2000/ants"
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

type Task struct {
	name string
	cmd  string
}

func exec_shell(t Task) {
	cmd := exec.Command("/bin/bash", "-c", t.cmd)
	//out := &myWriter{}
	log.SetFormatter(&MyJSONFormatter{})
	cmdLog := *log.WithFields(log.Fields{"name": t.name})
	cmd.Stdout = cmdLog.WriterLevel(log.InfoLevel)
	cmd.Stderr = cmdLog.WriterLevel(log.ErrorLevel)
	cmdLog.Infof("task %s start", t.name)
	err := cmd.Start()
	if err != nil {
		cmdLog.Error(err.Error())
		return
	}
	if err = cmd.Wait(); err != nil {
		cmdLog.Error(err.Error())
		return
	}
	cmdLog.Infof("task %s done", t.name)
	return
}

func Run() {
	runTimes := 10
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) error {
		exec_shell(i.(Task))
		wg.Done()
		return nil
	})

	tasks := make([]Task, runTimes)

	for i:=0; i<runTimes; i++ {
		tasks[i] = Task{
			name: fmt.Sprintf("task-%d", i),
			cmd: "/root/github/go/src/shell/t2/job.sh",
		}
	}

	for _, task := range tasks {
		wg.Add(1)
		p.Serve(task)
	}

	wg.Wait()
}

func main() {
	Run()
}