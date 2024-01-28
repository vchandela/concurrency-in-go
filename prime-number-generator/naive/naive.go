package naive

import (
	"cp/utils"
	"fmt"
	"time"
)

var (
	exit        = false
	prime_val   = 0
	found_prime = false
)

func print_prime() {
	for !exit {
		// Busy wait on found_prime
		for !found_prime && !exit {
			time.Sleep(100 * time.Millisecond)
		}
		if !exit {
			fmt.Println("Print: ", prime_val)
			prime_val = 0
			found_prime = false
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
		fmt.Println("Found: ", prime_val)
		found_prime = true
		for found_prime && !exit {
			time.Sleep(100 * time.Millisecond)
		}
		i += 1
	}

}

func Run() {
	go find_prime()
	go print_prime()
	time.Sleep(3 * time.Second)
	exit = true
}
