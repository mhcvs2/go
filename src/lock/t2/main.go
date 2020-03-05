package main

import "sync"

func t1() {
	var mailbox uint8
	var lock sync.RWMutex

	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())
}
