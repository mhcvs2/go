package main

import "designPattern/abstractFactory/human"

func main() {
	black := human.CreateSexHuman(human.FEMALEBLACKHUMAN)
	black.GetColor()
	black.Talk()
	black.GetSex()

	yellow := human.CreateSexHuman(human.MALEBLACKHUMAN)
	yellow.GetColor()
	yellow.Talk()
	yellow.GetSex()

}
