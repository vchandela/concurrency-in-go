package main

import (
	"sync"
	"testing"
)

func BenchmarkFileSearchST(b *testing.B) {
	fileSearchST("/Users/vivekchandela/personal", "README.md")
}

func BenchmarkFileSearchMT(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fileSearchMT("/Users/vivekchandela/personal", "README.md", wg)
	}(&wg)
	wg.Wait()
}
