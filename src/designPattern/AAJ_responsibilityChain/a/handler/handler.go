package handler

import "fmt"

type LevelType int

const (
	NULL LevelType = iota
	FATHER
	HUSBAND
	SON
)

type IBaseHandler interface {
	HandleMessage(wm IWoman)
	SetNextHandler(handler IBaseHandler)
}

type IHandler interface {
	GetLevel() LevelType
	Response(wm IWoman)
}

type Handler struct {
	handler IHandler
	nextHandler IBaseHandler
}

func NewHandler(handler IHandler) *Handler {
	return &Handler{handler:handler}
}

func (c *Handler) HandleMessage(wm IWoman) {
	if wm.GetType() == c.handler.GetLevel() {
		c.handler.Response(wm)
	} else if c.nextHandler != nil {
		c.nextHandler.HandleMessage(wm)
	} else {
		fmt.Println("no handler to request")
	}
}

func (c *Handler) SetNextHandler(handler IBaseHandler) {
	c.nextHandler = handler
}

/////////////////////////////////////////////////////////////////

type Father struct {
	Handler
	level LevelType
}

func NewFather() *Father {
	return &Father{level: FATHER}
}

func (h *Father) Response(wm IWoman) {
	fmt.Println("----daughter request father-----")
	fmt.Println(wm.GetRequest())
	fmt.Println("father agree")
}

func (h *Father) GetLevel() LevelType {
	return h.level
}

/////////////////////////////////////////////////////////////////

type Husband struct {
	Handler
	level LevelType
}

func NewHusband() *Husband {
	return &Husband{level: FATHER}
}

func (h *Husband) Response(wm IWoman) {
	fmt.Println("----wife request Husband-----")
	fmt.Println(wm.GetRequest())
	fmt.Println("Husband agree")
}

func (h *Husband) GetLevel() LevelType {
	return h.level
}
/////////////////////////////////////////////////////////////////

type Son struct {
	Handler
	level LevelType
}

func NewSon() *Son {
	return &Son{level: SON}
}

func (h *Son) Response(wm IWoman) {
	fmt.Println("----mother request Son-----")
	fmt.Println(wm.GetRequest())
	fmt.Println("Son agree")
}

func (h *Son) GetLevel() LevelType {
	return h.level
}