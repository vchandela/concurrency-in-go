package utils

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func worker2(wg *sync.WaitGroup, id int, jobs <-chan int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
	}
}

func WaitUsingWaitGroup() {
	jobs := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, id int, jobs <-chan int) {
			worker2(wg, id, jobs)
		}(&wg, w, jobs)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//Wait for all threads to complete.
	wg.Wait()
}
