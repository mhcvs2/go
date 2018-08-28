package handler

import "fmt"

type IProcessor interface {
	Process(wm IWoman)
}

type Processor struct {
	Handlers []IBaseHandler
}

func NewProcessor(handlers []IBaseHandler) *Processor {
	return &Processor{Handlers:handlers}
}

func (p *Processor) Process(wm IWoman) {
	for _, handler := range p.Handlers {
		if handler.Handles(wm) {
			handler.HandleMessage(wm)
			return
		}
	}
	fmt.Println("no hander to handle")
}