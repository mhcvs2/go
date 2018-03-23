package office

import "designPattern/facade/letterProcess"

type IOffice interface {
	SendLetter(context string, address string)
}

type ModenPostOffice struct {
	lp letterProcess.ILetterProcess
}

func NewModenPostOffice() *ModenPostOffice {
	return &ModenPostOffice{
		lp:new(letterProcess.Imp1),
	}
}

func (o *ModenPostOffice)SendLetter(context string, address string) {
	o.lp.Write(context)
	o.lp.Fill(address)
	o.lp.Send()
}