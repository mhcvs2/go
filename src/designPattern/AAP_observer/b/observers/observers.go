package observers

import (
	"fmt"
	"designPattern/AAP_observer/b/observer"
)

type Lisi struct {}
type Wangsi struct {}

func (l *Lisi) Update(o observer.IObservable, arg interface{}) {
	fmt.Println("li si observe han fei zi")
	fmt.Println("bao gao:", arg.(string))
}

func (l *Wangsi) Update(o observer.IObservable, arg interface{}) {
	fmt.Println("wang si observe han fei zi")
	fmt.Println("wo kan jian:", arg.(string))
}