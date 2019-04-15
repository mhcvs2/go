package main

import (
	"fmt"
	"generate/stringer/painkiller"
)

func t1() {
	fmt.Println(painkiller.Paracetamol.String())
	fmt.Println(painkiller.Aspirin.String())

}

func main() {
	t1()
}
