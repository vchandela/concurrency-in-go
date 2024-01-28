package sem

import (
	"cp/utils"
	"fmt"
	"time"
)

var (
	exit      bool
	prime_val int
	sem_find  Semaphore
	sem_print Semaphore
)

func print_prime() {
	sem_find.Acquire()
	for !exit {
		// Wait for prime number to be available
		sem_find.Acquire()

		fmt.Println("Print: ", prime_val)
		prime_val = 0

		// Let the producer thread know that the current prime number is printed and it can find the next prime number
		sem_print.Release()
	}
}

func find_prime() {
	i := 1
	sem_print.Acquire()
	for !exit {
		for !utils.Is_prime(i) {
			i += 1
		}
		prime_val = i
		fmt.Println("Found: ", prime_val)

		// Let consumer thread know we have a prime number available for printing
		sem_find.Release()

		// Wait for consumer thread to print the prime number
		sem_print.Acquire()

		i += 1
	}

}

func Run() {
	sem_find = NewSemaphore(1)
	sem_print = NewSemaphore(1)
	exit = false
	prime_val = 0
	go find_prime()
	go print_prime()
	time.Sleep(1 * time.Millisecond)
	exit = true
}
