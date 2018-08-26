package colleague

import (
	"fmt"
)

type ActionType int

const (
	BUYCOMPUTER ActionType = iota
	SELLCOMPUTER
	OFFSELL
	CLEARSTOCK
)

type IMediator interface {
	Bind(objects... interface{})
	Execute(action ActionType, args... interface{})
}

//----------------------------------------
type Mediator struct {
	purchase *Purchase
	sale *Sale
	stock *Stock
	bindMap map[string]string
}

func NewMediator() *Mediator {
	return new(Mediator)
}

func (m *Mediator) Bind(objects... interface{}) {

	for _, ob := range objects {
		tname := GetTypeName(ob)
		switch tname {
		case "Purchase":
			m.purchase = ob.(*Purchase)
		case "Sale":
			m.sale = ob.(*Sale)
		case "Stock":
			m.stock = ob.(*Stock)
		default:
			fmt.Println("unknow type", tname)
		}
	}
}

func (m *Mediator)Execute(action ActionType, args... interface{}) {
	switch action {
	case BUYCOMPUTER:
		m.buyComputer(args[0].(int))
	case SELLCOMPUTER:
		m.SellComputer(args[0].(int))
	case OFFSELL:
		m.OffSell()
	case CLEARSTOCK:
		m.ClearStock()
	default:
		fmt.Println("unknow action", action)
	}
}

func (m *Mediator)buyComputer(number int) {
	saleStatus := m.sale.GetSaleStatus()
	var buyNumber int
	if saleStatus <= 80 {
		buyNumber = number / 2
	} else {
		buyNumber = number
	}
	fmt.Println("buy IBM computer number:", buyNumber)
	m.stock.Increase(buyNumber)
}

func (m *Mediator) SellComputer(number int) {
	if m.stock.GetStockNumber() < number {
		m.purchase.BuyIBMComputer(number)
	}
	m.stock.Decrease(number)
}

func (m *Mediator) OffSell() {
	fmt.Println("off sell number:", m.stock.GetStockNumber())
}

func (m *Mediator) ClearStock() {
	m.sale.OffSale()
	m.purchase.RefuseBuyIBM()
}