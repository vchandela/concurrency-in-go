package utils

import (
	"fmt"
	"sync"
)

type Barrier struct {
	total int
	count int
	mutex *sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(size int) *Barrier {
	lock := &sync.Mutex{}
	condVar := sync.NewCond(lock)
	return &Barrier{
		total: size,
		count: size,
		mutex: lock,
		cond:  condVar,
	}
}

func (b *Barrier) Wait(threadName string) {
	b.mutex.Lock()
	b.count -= 1

	// we don't use for-loop here as broadcasting thread is setting b.count = b.total
	if b.count > 0 {
		fmt.Println(threadName, "waiting on condition")
		// unlocks mutex and suspend execution of current goroutine
		b.cond.Wait()
	} else {
		fmt.Println(threadName, "releases all waiting threads")
		b.cond.Broadcast()
		b.count = b.total
	}

	b.mutex.Unlock()
}
