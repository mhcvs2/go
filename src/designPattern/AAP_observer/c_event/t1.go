package main

import (
	"designPattern/AAP_observer/c_event/t1"
	"time"
)

func main() {
	hanfeizi := t1.NewHanFeiZi()
	hanfeizi.HaveBreakfast()
	hanfeizi.HaveFun()
	time.Sleep(time.Second * 5)
	hanfeizi.GetEventManager().Stop(true)
}
