package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

var (
	frequencyST       [26]int32
	lock              = sync.Mutex{}
	frequencyMTLock   [26]int32
	frequencyMTAtomic [26]int32
	wg                sync.WaitGroup
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Single-threaded
func countLettersST(url string, frequency *[26]int32) {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index >= 0 {
			frequency[index] += 1
		}
	}
}

// Multi-threaded using lock
func countLettersMTLock(wg *sync.WaitGroup, url string, frequency *[26]int32) {
	defer wg.Done()
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index >= 0 {
			lock.Lock()
			frequency[index] += 1
			lock.Unlock()
		}
	}
}

// Multi-threaded using atomic
func countLettersMTAtomic(wg *sync.WaitGroup, url string, frequency *[26]int32) {
	defer wg.Done()
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		index := strings.Index(allLetters, c)
		if index >= 0 {
			atomic.AddInt32(&frequency[index], 1)
		}
	}
}

func main() {
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		countLettersST(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequencyST)
	}
	elapsed := time.Since(start)
	fmt.Println("Single-threaded Elapsed time:", elapsed)
	// for i, f := range frequencyST {
	// 	fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	// }

	start = time.Now()
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go countLettersMTLock(&wg, fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequencyMTLock)
	}
	wg.Wait()
	elapsed = time.Since(start)
	fmt.Println("Multi-threaded with Lock Elapsed time:", elapsed)
	// for i, f := range frequencyMTLock {
	// 	fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	// }

	start = time.Now()
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go countLettersMTAtomic(&wg, fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequencyMTAtomic)
	}
	wg.Wait()
	elapsed = time.Since(start)
	fmt.Println("Multi-threaded with Atomic Elapsed time:", elapsed)
	// for i, f := range frequencyMTAtomic {
	// 	fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	// }
}
