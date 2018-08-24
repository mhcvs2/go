package main

import "designPattern/AAD_template/hanma"

func main() {
	h1 := hanma.HanmaH1Model{}
	h2 := hanma.HanmaH2Model{}
	h := hanma.HanmaModel{Specific:h1}
	h.Run()
	h.Specific = h2
	h.Run()
}