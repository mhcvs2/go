package model

type Trade struct {
	TradeNo string
	Amount int
}

func NewTrade() *Trade {
	return &Trade{}
}
