package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync/atomic"
)

type job func()

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	jobID uint32 = 0
)

func getJobs() []job {
	res := make([]job, 10)
	for i:=0; i<10; i++ {
		res[i] = func() {
			atomic.AddUint32(&jobID, 1)
			id := jobID
			rd := r.Intn(10)
			str := fmt.Sprintf("job-%d done, time: %ds", id, rd)
			time.Sleep(time.Second * time.Duration(rd))
			fmt.Println(str)
		}
	}
	return res
}

func main() {
	jobList := getJobs()

	jobs := make(chan job)
	done := make(chan bool, len(jobList))

	go func() {
		for _, job := range jobList {
			jobs <- job
		}
		close(jobs)
	}()

	go func() {
		for job := range jobs {
			go func() {
				job()
				done <- true
			}()
		}
	}()

	for i:=0; i<len(jobList); i++ {
		<-done
	}
	fmt.Println("all done, exit")
}

//job-2 done, time: 0s
//job-3 done, time: 0s
//job-9 done, time: 0s
//job-1 done, time: 2s
//job-4 done, time: 2s
//job-5 done, time: 2s
//job-7 done, time: 2s
//job-10 done, time: 6s
//job-6 done, time: 8s
//job-8 done, time: 9s
//all done, exit