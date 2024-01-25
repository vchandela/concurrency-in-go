package utils

import "testing"

func BenchmarkWaitUsingChannel(b *testing.B) {
	WaitUsingChannel()
}

func BenchmarkWaitUsingWaitGroup(b *testing.B) {
	WaitUsingWaitGroup()
}
