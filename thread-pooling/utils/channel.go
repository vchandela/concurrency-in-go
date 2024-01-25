package utils

import (
	"fmt"
	"time"
)

const (
	numJobs = 5
)

func worker(id int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		res <- j * 2
	}
}

func WaitUsingChannel() {
	jobs := make(chan int, numJobs)
	res := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, res)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		<-res
	}
	close(res)
}
