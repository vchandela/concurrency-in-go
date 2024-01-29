package cond_var

import (
	"bb/queue"
	"fmt"
	"sync"
)

type BlockingQueue interface {
	Enqueue(threadName string, val int)
	Dequeue(threadName string) int
}

type blockingQueue struct {
	maxxSize int
	cond     *sync.Cond
	q        queue.Queue
}

func NewBlockingQueue(maxxSize int) BlockingQueue {
	lock := sync.Mutex{}
	cond := sync.NewCond(&lock)
	q := queue.NewQueue()
	return &blockingQueue{
		maxxSize, cond, q,
	}
}

func (bq *blockingQueue) Enqueue(threadName string, val int) {
	bq.cond.L.Lock()
	for bq.q.Len() == bq.maxxSize {
		bq.cond.Wait()
	}
	bq.q.Push(val)
	fmt.Printf("Thread %v produced new val %v; queue size is %v\n", threadName, val, bq.q.Len())
	bq.cond.Broadcast()
	bq.cond.L.Unlock()
}

func (bq *blockingQueue) Dequeue(threadName string) int {
	bq.cond.L.Lock()
	for bq.q.Len() == 0 {
		bq.cond.Wait()
	}
	val := bq.q.Pop()
	fmt.Printf("Thread %v consumer val %v; queue size is %v\n", threadName, val, bq.q.Len())
	bq.cond.Broadcast()
	bq.cond.L.Unlock()
	return val.(int)
}
