package main

import (
	"designPattern/ABJ_factory_strategy/deduction"
	"designPattern/ABJ_factory_strategy/model"
	"fmt"
)

func initIC() *model.Card {
	card := model.NewCard()
	card.SteadyMoney = 800
	card.FreeMoney = 1000
	card.CardNo = "101020030101"
	return card
}

func createTrade() *model.Trade {
	trade := model.NewTrade()
	trade.TradeNo = "sdfabc"
	trade.Amount = 200
	return trade
}

func main() {
	card := initIC()
	trade := createTrade()
	deduction.Deduct(card, trade)
	fmt.Println(card.FreeMoney)
	fmt.Println(card.SteadyMoney)
}
