package observer

import "fmt"

type Obr interface {
	Update(context string)
}

type Lisi struct {}
type Wangsi struct {}

func (l *Lisi) Update(context string) {
	fmt.Println("li si observe han fei zi")
	fmt.Println("bao gao:", context)
}

func (l *Wangsi) Update(context string) {
	fmt.Println("wang si observe han fei zi")
	fmt.Println("wo kan jian:", context)
}