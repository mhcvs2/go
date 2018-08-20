package main

import "designPattern/AAB_factory/human"

func main() {
	black := human.CreateHuman(human.BLACKHUMAN)
	black.GetColor()
	black.Talk()

	yellow := human.CreateHuman(human.YELLOWHUMAN)
	yellow.GetColor()
	yellow.Talk()

}
