package deduction

import (
	"designPattern/ABJ_factory_strategy/model"
	"math"
)

type IDeduction interface {
	Exec(card *model.Card, trade *model.Trade) bool
}

type FreeDeduction struct {}

func (d *FreeDeduction) Exec(card *model.Card, trade *model.Trade) bool {
	card.FreeMoney -= trade.Amount
	return true
}

type SteadyDeduction struct {}

func (d *SteadyDeduction) Exec(card *model.Card, trade *model.Trade) bool {
	halfMoney := int(math.Floor(float64(trade.Amount) / 2))
	card.FreeMoney -= halfMoney
	card.SteadyMoney -= halfMoney
	return true
}

