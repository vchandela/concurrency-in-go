package cond_var

import (
	"cp/utils"
	"fmt"
	"sync"
	"time"
)

var (
	exit        bool
	prime_val   int
	found_prime bool
	lock        sync.Mutex
	cond_var    *sync.Cond
)

func print_prime() {
	for !exit {
		cond_var.L.Lock()
		// Once thread wakes up, while-loop makes sure found_prime is true to move forward
		for !found_prime && !exit {
			cond_var.Wait()
		}
		cond_var.L.Unlock()

		if !exit {
			fmt.Println("Print: ", prime_val)
			prime_val = 0

			cond_var.L.Lock()
			found_prime = false
			cond_var.Signal()
			cond_var.L.Unlock()
		}
	}
}

func find_prime() {
	i := 1
	for !exit {
		for !utils.Is_prime(i) {
			i += 1
		}
		prime_val = i

		cond_var.L.Lock()
		found_prime = true
		fmt.Println("Found: ", prime_val)
		cond_var.Signal()
		cond_var.L.Unlock()

		cond_var.L.Lock()
		//  Once thread wakes up, while-loop makes sure found_prime is false to move forward
		for found_prime && !exit {
			cond_var.Wait()
		}
		cond_var.L.Unlock()

		i += 1
	}

}

func Run() {
	cond_var = sync.NewCond(&lock)
	exit = false
	prime_val = 0
	found_prime = false
	go find_prime()
	go print_prime()
	time.Sleep(3 * time.Millisecond)
	exit = true
}
