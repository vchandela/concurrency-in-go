package main

import (
	"fmt"
	"testing"
)

func BenchmarkST(b *testing.B) {
	for i := 1000; i <= 1200; i++ {
		countLettersST(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequencyST)
	}
}

func BenchmarkMTLock(b *testing.B) {
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go countLettersMTLock(&wg, fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequencyMTLock)
	}
}

func BenchmarkMTAtomic(b *testing.B) {
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go countLettersMTAtomic(&wg, fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequencyMTAtomic)
	}
}
