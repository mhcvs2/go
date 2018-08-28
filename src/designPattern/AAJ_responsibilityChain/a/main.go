package main

import "designPattern/AAJ_responsibilityChain/a/handler"

func main() {
	f := handler.NewHandler(handler.NewFather())
	h := handler.NewHandler(handler.NewHusband())
	s := handler.NewHandler(handler.NewSon())
	f.SetNextHandler(h)
	h.SetNextHandler(s)

	wm := handler.NewWoman(handler.FATHER, "playing")
	f.HandleMessage(wm)

	wm = handler.NewWoman(handler.HUSBAND, "shopping")
	f.HandleMessage(wm)

	wm = handler.NewWoman(handler.SON, "shangfen")
	f.HandleMessage(wm)
}

//----daughter request father-----
//daughter's request is: playing
//father agree
//----wife request Husband-----
//wife's request is: shopping
//Husband agree
//----mother request Son-----
//mother's request is: shangfen
//Son agree