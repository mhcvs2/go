package main

import (
	"runtime"
	"regexp"
	"time"
	"fmt"
	"os"
	"log"
	"bufio"
	"bytes"
	"io"
	"github.com/mhcvs2/godatastructure/util"
	"path/filepath"
)

type Result struct {
	filename string
	lino int
	line string
}

type Job struct {
	filename string
	results chan<- Result
}

func (j *Job) Do(lineRx *regexp.Regexp) {
	file, err := os.Open(j.filename)
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n\r")
		if lineRx.Match(line) {
			j.results <- Result{j.filename, lino, string(line)}
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("error: %d: %s\n", lino, err)
			}
			break
		}
	}
}

//---------------------------------------------------
var workers = runtime.NumCPU()

func addJobs(jobs chan<- Job, filenames []string, results chan<- Result) {
	for _, filename := range filenames {
		jobs <- Job{filename:filename, results:results}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, lineRx *regexp.Regexp, jobs <-chan Job) {
	for job := range jobs {
		job.Do(lineRx)
	}
	done <- struct{}{}
}

func waitAndProcessResults(timeout int64, done <-chan struct{}, results <-chan Result) {
	finish := time.After(time.Second * time.Duration(timeout))
	for working := workers; working > 0; {
		select {
		case result := <-results:
			fmt.Printf("%s: line-%d: %s\n", filepath.Base(result.filename), result.lino, result.line)
		case <-finish:
			fmt.Println("timeout")
			return
		case <-done:
			working--
		}
	}
	for{
		select {
		case result := <- results:
			fmt.Printf("%s: %d: %s\n", filepath.Base(result.filename), result.lino, result.line)
		case <-finish:
			fmt.Println("timeout")
			return
		default:
			return
		}
	}
}

func grep(lineRx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, util.Minimum(1000, len(filenames)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, filenames, results)

	for i:=0; i<workers; i++ {
		go doJobs(done, lineRx, jobs)
	}

	waitAndProcessResults(10, done, results)
}

//------------------------------------------------------------
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rege := "mhc"
	files, err := util.GetSubFiles("/root/github/go/src/chan/cgrep/testfiles")
	if err != nil {
		log.Fatalf("list files error: %s\n", err.Error())
	}
	lineRx, err := regexp.Compile(rege)
	if err != nil {
		log.Fatalf("invalid regext: %s\n", rege)
	}
	grep(lineRx, files)
}
