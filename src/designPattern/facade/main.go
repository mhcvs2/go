package main

import (
	"designPattern/facade/office"
)

func main() {
	o := office.NewModenPostOffice()
	o.SendLetter("hello jack", "beijing")
}

//write: hello jack
//fill address: beijing
//send letter