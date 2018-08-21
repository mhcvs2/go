package human

import "fmt"

type Human interface {
	GetColor()
	Talk()
}

type SexHuman interface {
	Human
	GetSex()
}
//----------------------------------
type BlackHuman struct {}

func (h *BlackHuman) GetColor() {
	fmt.Println("BlackHuman GetColor")
}

func (h *BlackHuman) Talk() {
	fmt.Println("BlackHuman Talk")
}

type FemaleBlackHuman struct {
	BlackHuman
}

func (h *FemaleBlackHuman)GetSex() {
	fmt.Println("Female Black Human")
}

type MaleBlackHuman struct {
	BlackHuman
}

func (h *MaleBlackHuman)GetSex() {
	fmt.Println("Male Black Human")
}
//----------------------------------
type YellowHuman struct {}

func (h YellowHuman) GetColor() {
	fmt.Println("YellowHuman GetColor")
}

func (h YellowHuman) Talk() {
	fmt.Println("YellowHuman Talk")
}
//----------------------------------
type WhiteHuman struct {}

func (h WhiteHuman) GetColor() {
	fmt.Println("WhiteHuman GetColor")
}

func (h WhiteHuman) Talk() {
	fmt.Println("WhiteHuman Talk")
}
