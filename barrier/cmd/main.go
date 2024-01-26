package main

import (
	"barrier/utils"
	"fmt"
	"time"
)

func waitOnBarrier(threadName string, timeToSleep int, barrier *utils.Barrier) {
	for {
		fmt.Println(threadName, "running")
		time.Sleep(time.Duration(timeToSleep) * time.Second)
		barrier.Wait(threadName)
	}
}

func main() {
	barrier := utils.NewBarrier(3)
	go waitOnBarrier("red", 4, barrier)
	go waitOnBarrier("blue", 10, barrier)
	go waitOnBarrier("green", 6, barrier)
	time.Sleep(time.Duration(50) * time.Second)
}
