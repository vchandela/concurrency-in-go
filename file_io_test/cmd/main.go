package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"file_io_test/utils"

	"github.com/sourcegraph/conc/pool"
)

type txtLine struct {
	id  int
	txt string
}

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:n], nil
}

func generate_string() {
	file, err := os.Create("../temp100K.txt")
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	defer file.Close()

	for i := 0; i < 100000; i++ {
		randomString, err := GenerateRandomString(16) // Each string is 16 characters long
		if err != nil {
			log.Fatalf("failed to generate random string: %s", err)
		}

		_, err = file.WriteString(randomString + "\n")
		if err != nil {
			log.Fatalf("failed to write to file: %s", err)
		}
	}

	fmt.Println("Random strings generated and written to temp100K.txt")
}

func pollRecords(reader *bufio.Reader, max_poll_records int) []txtLine {
	lines := make([]txtLine, 0)
	for i := 0; i < max_poll_records; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" && len(line) > 0 {
				lines = append(lines, txtLine{i, line})
			}
			break
		}
		lines = append(lines, txtLine{i, line})
	}
	return lines
}

func ioTest(max_poll_records, channel_size, num_workers int) {
	// Get the current number of CPUs
	// numCPU := runtime.NumCPU()
	// fmt.Println("Number of CPUs:", numCPU)

	// Set GOMAXPROCS to the desired value
	runtime.GOMAXPROCS(4)
	// fmt.Println("GOMAXPROCS is set to:", runtime.GOMAXPROCS(0)) // Passing 0 returns the current value

	tasks := make(chan txtLine, channel_size)

	var workers = num_workers
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			for range tasks {
				// fmt.Printf("worker %d processed record %d\n", id, record.id)
				time.Sleep(2 * time.Millisecond)
			}
		}(i, &wg)
	}

	file, err := os.Open("../temp100K.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		fetches := pollRecords(reader, max_poll_records)
		// fmt.Println("fetched ", len(fetches), " records")
		if len(fetches) == 0 {
			close(tasks)
			wg.Wait()
			break
		}
		for _, record := range fetches {
			// fmt.Printf("stored id %d, val %s\n", record.id, record.txt)
			tasks <- record
		}
	}

}

func ioTest2(max_poll_records int) {
	// Set GOMAXPROCS to the desired value
	runtime.GOMAXPROCS(4)

	concurrency := utils.GetCurrentProcessConc()
	fmt.Println("Concurrency:", concurrency)
	pl := pool.New().WithMaxGoroutines(concurrency)

	file, err := os.Open("../temp100K.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		fetches := pollRecords(reader, max_poll_records)
		
		if len(fetches) == 0 {
			break
		}
		for range fetches {
			pl.Go(func() {
				time.Sleep(2 * time.Millisecond)
			})
		}
	}

}

func ioTest3(max_poll_records, num_workers int) {
	// Set GOMAXPROCS to the desired value
	runtime.GOMAXPROCS(4)

	pl := pool.New().WithMaxGoroutines(num_workers)

	file, err := os.Open("../temp100K.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		fetches := pollRecords(reader, max_poll_records)
		if len(fetches) == 0 {
			break
		}
		for range fetches {
			pl.Go(func() {
				time.Sleep(2 * time.Millisecond)
			})
		}
	}

}

func main() {
	// ioTest(constants.MAX_POLL_RECORDS, constants.MAX_POLL_RECORDS, constants.NUM_WORKERS)
	generate_string()
}
