package main

import "designPattern/AAC_abstractFactory/a/human"

func main() {
	black := new(human.MaleFactory).CreateBlackHuman()
	black.GetColor()
	black.Talk()
	black.GetSex()

	yellow := new(human.FemaleFactory).CreateBlackHuman()
	yellow.GetColor()
	yellow.Talk()
	yellow.GetSex()

}
