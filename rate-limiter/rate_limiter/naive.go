package rate_limiter

import (
	"fmt"
	"time"
)

func Naive() {
	// incoming requests go to a channel
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Limit handler to 1 request every 200 ms
	limiter := time.NewTicker(200 * time.Millisecond)
	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}
	limiter.Stop()
}
