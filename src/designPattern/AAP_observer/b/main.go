package main

import (
	"designPattern/AAP_observer/b/observers"
	"designPattern/AAP_observer/b/observable"
)

func main() {
	h := observable.NewHanFeiZi()

	l := new(observers.Lisi)
	w := new(observers.Wangsi)
	h.AddObserver(l)
	h.AddObserver(w)
	h.HaveBreakfast()
	h.HaveFun()
}

//i bave breakfast
//li si observe han fei zi
//bao gao: han fei zi have breakfast
//wang si observe han fei zi
//wo kan jian: han fei zi have breakfast
//i have fun
//li si observe han fei zi
//bao gao: han fei zi have fun
//wang si observe han fei zi
//wo kan jian: han fei zi have fun