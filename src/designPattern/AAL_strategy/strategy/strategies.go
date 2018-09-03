package strategy

import "fmt"

type IStrategy interface {
	operate()
}

type BackDoor struct {}

func (s *BackDoor) operate() {
	fmt.Println("back door...")
}

type GivenGreenLight struct {}

func (s *GivenGreenLight) operate() {
	fmt.Println("given green light")
}

type BlockEnemy struct {}

func (s *BlockEnemy) operate() {
	fmt.Println("block enemy")
}