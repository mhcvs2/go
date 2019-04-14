package model

type Card struct {
	CardNo string
	SteadyMoney int
	FreeMoney int
}

func NewCard() *Card {
	return &Card{}
}


