package main

import (
	"fmt"
	"designPattern/AAY_factory_builder/c_builder_superman/superman"
)

func main() {
	adultSuperMan := superman.GetAdultSuperMan()
	fmt.Println(adultSuperMan.SpecialTalent)
}
