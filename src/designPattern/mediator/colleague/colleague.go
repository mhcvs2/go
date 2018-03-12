package colleague

import (
	"math/rand"
	"time"
	"fmt"
)

type Colleague struct {
	mediator IMediator
}

//-----------------------------------
type Purchase struct {
	*Colleague
}

func NewPurchase(mediator IMediator) *Purchase {
	p := &Purchase{&Colleague{mediator:mediator}}
	mediator.Bind(p)
	return p
}

func (p *Purchase) BuyIBMComputer(number int) {
	p.mediator.Execute(BUYCOMPUTER, number)
}

func (p *Purchase) RefuseBuyIBM(){
	fmt.Println("not buy IBM")
}
//-----------------------------------
type Stock struct {
	*Colleague
	couputerNumber int
}

func NewStock(mediator IMediator) *Stock {
	s := &Stock{
		Colleague:&Colleague{mediator:mediator},
		couputerNumber:100,
	}
	mediator.Bind(s)
	return s
}

func (s *Stock) Increase(number int) {
	s.couputerNumber += number
	fmt.Println("stock number is:", s.couputerNumber)
}

func (s *Stock) Decrease(number int) {
	s.couputerNumber -= number
	fmt.Println("stock number is:", s.couputerNumber)
}

func (s *Stock) GetStockNumber() int {
	return s.couputerNumber
}

func (s *Stock) ClearStock() {
	fmt.Println("clear stock number:", s.couputerNumber)
	s.mediator.Execute(CLEARSTOCK)
}

//-----------------------------------
type Sale struct {
	*Colleague
}

func NewSale(mediator IMediator) *Sale {
	s := &Sale{&Colleague{mediator:mediator}}
	mediator.Bind(s)
	return s
}

func (s *Sale) SellIBMComputer(number int) {
	s.mediator.Execute(SELLCOMPUTER, number)
	fmt.Println("sell commputer number: ", number)
}

func (s *Sale) GetSaleStatus() int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	saleStatus := rd.Intn(100)
	fmt.Println("sale status is:", saleStatus)
	return saleStatus
}

func (s *Sale) OffSale(){
	s.mediator.Execute(OFFSELL)
}

