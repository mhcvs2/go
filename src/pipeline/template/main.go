package main

import (
	"time"
	"pipeline/template/pipline"
	"fmt"
)

type Test struct {
}

func (t Test)T(a string) string {
	time.Sleep(time.Second * 3)
	return a + " done"
}

func t1() {
	done := make(chan struct{})
	defer close(done)
	t := Test{}
	p := pipline.Pipline{Ob:t, Action:"T", Args:[]string{"hi11","hi22"}, Wn:3, Done:done}
	res := p.Run()
	fmt.Println(len(res))
	for _, r := range res{
		fmt.Println(r)
	}
}

func t2() {
	done := make(chan struct{})
	defer close(done)
	t := Test{}
	p := pipline.NewPipline(t, "T", []string{"hi11","hi22"}, done)
	p.Run()
}

func t3() {
	done := make(chan struct{})
	defer close(done)
	t := Test{}
	res := pipline.RunPipline(t, "T", []string{"hi11","hi22"}, done)
	for _, r := range res{
		fmt.Println(r)
	}
}

func main() {
	t3()
}