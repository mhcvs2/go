package letterProcess

import "fmt"

type ILetterProcess interface {
	Write(context string)
	Fill(address string)
	Send()
}

type Imp1 struct {}

func (l *Imp1) Write(context string) {
	fmt.Println("write:", context)
}

func (l *Imp1) Fill(address string) {
	fmt.Println("fill address:", address)
}

func (l *Imp1) Send() {
	fmt.Println("send letter")
}