package main

import (
	"os"
	"syscall"
	gosignal "os/signal"
	"github.com/cloud-pi/infra_building_backend/module/log"
	"sync/atomic"
	"time"
)

func Trap(cleanup func()) {
	c := make(chan os.Signal, 1)
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM}
	if os.Getenv("debug") == "" {
		signals = append(signals, syscall.SIGQUIT)
	}
	gosignal.Notify(c, signals...)
	go func() {
		count := uint32(0)
		for sig := range c {
			go func(sig os.Signal) {
				log.Printf("Received signal '%v', starting shutdown...\n", sig)
				switch sig {
				case os.Interrupt, syscall.SIGTERM:
					if atomic.LoadUint32(&count) < 3 {
						atomic.AddUint32(&count, 1)
						if atomic.LoadUint32(&count) == 1 {
							cleanup()
							os.Exit(0)
						} else {
							return
						}
					} else {
						log.Printf("Force shutdown...\n")
					}
				case syscall.SIGQUIT:
				}
				os.Exit(128 + int(sig.(syscall.Signal)))
			}(sig)
		}
	}()
}

//--------------------------------------------------------

type Some struct {
	A string
}

func NewSome(a string) *Some {
	return &Some{A: a}
}

func (s *Some)CleanUp() {
	log.Println("I am some, i will quit...")
	log.Printf("A: %s\n", s.A)
	time.Sleep(time.Second * 10)
	log.Println("done")
}

//-----------------------------------------------------------

func main() {
	s := NewSome("haha")
	Trap(s.CleanUp)
	work := make(chan string)
	for i := range work {
		log.Println(i)
	}
}