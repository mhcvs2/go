package emperor2

import (
	"fmt"
	"sync"
)

type emperor struct {
	Name string
}

func (e emperor) Say() {
	fmt.Printf("my name: %s\n", e.Name)
}

var singleton *emperor
var lock *sync.Mutex = &sync.Mutex{}

func GetInstance() *emperor {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			singleton = &emperor{}
		}
	}
	return singleton
}
