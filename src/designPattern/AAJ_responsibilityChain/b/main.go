package main

import "designPattern/AAJ_responsibilityChain/b/handler"

func main() {
	f := handler.NewHandler(handler.NewFather())
	h := handler.NewHandler(handler.NewHusband())
	s := handler.NewHandler(handler.NewSon())
	processor := handler.NewProcessor([]handler.IBaseHandler{f, h, s})

	wm := handler.NewWoman(handler.FATHER, "playing")
	processor.Process(wm)

	wm = handler.NewWoman(handler.HUSBAND, "shopping")
	processor.Process(wm)

	wm = handler.NewWoman(handler.SON, "shangfen")
	processor.Process(wm)
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