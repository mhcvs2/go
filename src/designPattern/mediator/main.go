package main

import "designPattern/mediator/colleague"

func main() {
	mediator := colleague.NewMediator()
	p := colleague.NewPurchase(mediator)
	s1 := colleague.NewSale(mediator)
	s2 := colleague.NewStock(mediator)

	p.BuyIBMComputer(100)
	s1.SellIBMComputer(1)
	s2.ClearStock()
}
