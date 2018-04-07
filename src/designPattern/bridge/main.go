package main

import "fmt"

type IProduct interface {
	BeProducted()
	BeSelled()
}

//---------------------------------
type House struct {}

func (h *House) BeProducted() {
	fmt.Println("product house")
}

func (h *House) BeSelled() {
	fmt.Println("sell house")
}

//---------------------------------
type IPod struct {}

func (h *IPod) BeProducted() {
	fmt.Println("product IPod")
}

func (h *IPod) BeSelled() {
	fmt.Println("sell IPod")
}

//---------------------------------
type ICorp interface {
	MakeMoney()
}

type Corp struct {
	product IProduct
}

func NewCorp(prduct IProduct) *Corp {
	return &Corp{product:prduct}
}

func (c *Corp)MakeMoney() {
	c.product.BeProducted()
	c.product.BeSelled()
}

//-----------------------------------------
type HouseCorp struct {
	Corp
}

func NewHouseCorp(prduct *House) *HouseCorp {
	return &HouseCorp{Corp:*NewCorp(prduct)}
}

func (h *HouseCorp) MakeMoney() {
	h.Corp.MakeMoney()
	fmt.Println("House Corp make big money...")
}

//------------------------------------------
type ShanZhaiCorp struct {
	Corp
}

func NewShanZhaiCorp(prduct IProduct) *ShanZhaiCorp {
	return &ShanZhaiCorp{Corp:*NewCorp(prduct)}
}

func (h *ShanZhaiCorp) MakeMoney() {
	h.Corp.MakeMoney()
	fmt.Println("I make money...")
}

//-------------------------------------------------
func main() {
	house := new(House)
	houseCorp := NewHouseCorp(house)
	houseCorp.MakeMoney()

	shanzhai := NewShanZhaiCorp(new(IPod))
	shanzhai.MakeMoney()
}

//product house
//sell house
//House Corp make big money...
//product IPod
//sell IPod
//I make money...