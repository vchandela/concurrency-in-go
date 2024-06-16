package main

import (
	"bufio"
	"fmt"
	"file_io_test/constants"
	"os"
	"runtime"
	"sync"
	"time"
)

type txtLine struct {
	id  int
	txt string
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
	runtime.GOMAXPROCS(8)
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

	file, err := os.Open("../temp.txt")
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

func main() {
	ioTest(constants.MAX_POLL_RECORDS, constants.MAX_POLL_RECORDS, constants.NUM_WORKERS)
}
