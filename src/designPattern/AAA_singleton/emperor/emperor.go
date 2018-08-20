package emperor

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
var once sync.Once

func GetInstance() *emperor {
	once.Do(func() {
		singleton = &emperor{}
	})
	return singleton
}
