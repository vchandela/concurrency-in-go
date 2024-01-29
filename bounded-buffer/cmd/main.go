package main

import (
	"bb/channel"
	"time"
)

// Fast producer
func Produce(bq channel.BlockingQueue, threadName string, val int) {
	item := val
	for {
		bq.Enqueue(threadName, item)
		item += 1
		time.Sleep(100 * time.Millisecond)
	}
}

// Slow consumer
func Consume(bq channel.BlockingQueue, threadName string) {
	for {
		bq.Dequeue(threadName)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	bq := channel.NewBlockingQueue(5)
	go Consume(bq, "c1")
	go Consume(bq, "c2")
	go Produce(bq, "p1", 1)
	go Produce(bq, "p2", 100)

	time.Sleep(1 * time.Second)
}
