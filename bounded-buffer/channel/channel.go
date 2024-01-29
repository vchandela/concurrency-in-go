package channel

import (
	"fmt"
	"sync/atomic"
)

type BlockingQueue interface {
	Enqueue(threadName string, val int)
	Dequeue(threadName string) int
}

type blockingQueue struct {
	currSize uint32
	c        chan int
}

func NewBlockingQueue(maxxSize int) BlockingQueue {
	c := make(chan int, maxxSize)
	return &blockingQueue{
		0, c,
	}
}

func (bq *blockingQueue) Enqueue(threadName string, val int) {
	bq.c <- val
	// adding 1 atomically to avoid race conditions
	atomic.AddUint32(&bq.currSize, 1)
	fmt.Printf("Thread %v produced new val %v; queue size is %v\n", threadName, val, bq.currSize)
}

func (bq *blockingQueue) Dequeue(threadName string) int {
	val := <-bq.c
	// subtracting 1 atomically to avoid race conditions
	atomic.AddUint32(&bq.currSize, ^uint32(0))
	fmt.Printf("Thread %v consumed val %v; queue size is %v\n", threadName, val, bq.currSize)
	return val
}
