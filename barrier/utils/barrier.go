package utils

import (
	"fmt"
	"sync"
)

type Barrier interface {
	Wait(threadName string)
}

type barrier struct {
	total int
	count int
	cond  *sync.Cond
}

func NewBarrier(size int) Barrier {
	lock := &sync.Mutex{}
	condVar := sync.NewCond(lock)
	return &barrier{
		total: size,
		count: size,
		cond:  condVar,
	}
}

func (b *barrier) Wait(threadName string) {
	b.cond.L.Lock()

	b.count -= 1

	// b.count < b.total is needed otherwise line #46 will keep the threads waiting
	for b.count > 0 && b.count < b.total {
		fmt.Println(threadName, "waiting on condition")
		// unlocks mutex and suspend execution of current goroutine
		b.cond.Wait()
	}

	// if condition is needed here, otherwise green and red threads will also Broadcast() and set b.count = 0 after waking up.
	if b.count == 0 {
		fmt.Println(threadName, "releases all waiting threads")
		b.cond.Broadcast()
		b.count = b.total
	}

	b.cond.L.Unlock()
}
