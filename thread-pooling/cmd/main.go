package main

import (
	"fmt"
	"pooling/utils"
	"time"
)

func main() {
	start := time.Now()
	utils.WaitUsingChannel()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)

	start = time.Now()
	utils.WaitUsingWaitGroup()
	elapsed = time.Since(start)
	fmt.Println("Elapsed time:", elapsed)

}
