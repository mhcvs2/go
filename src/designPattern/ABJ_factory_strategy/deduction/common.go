package deduction

import (
	"designPattern/ABJ_factory_strategy/model"
	"strings"
)

type DeductionContext struct {
	deduction IDeduction
}

func NewDeductionContext(deduction IDeduction) *DeductionContext {
	return &DeductionContext{deduction: deduction}
}

func (s *DeductionContext)Exec(card *model.Card, trade *model.Trade) bool {
    return s.deduction.Exec(card, trade)
}


// facade
func getDeductionType(trade *model.Trade) StragegyMan {
	if strings.Contains(trade.TradeNo, "abc") {
		return FREE_DEDUCTION
	}
	return STEADY_DEDUCTION
}

func Deduct(card *model.Card, trade *model.Trade) *model.Card {
	reg := getDeductionType(trade)
	deduction := GetDeduction(reg)
	context := NewDeductionContext(deduction)
	context.Exec(card, trade)
	return card
}
