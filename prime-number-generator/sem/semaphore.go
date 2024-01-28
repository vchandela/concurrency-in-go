package sem

// Interface abstracts the struct from outside access.
type Semaphore interface {
	Acquire()
	Release()
}

// The struct{} type has zero size, which means it doesn't consume any memory.
// This is useful when you don't need any data associated with the channel element and you're using the channel purely for signaling purposes.
type semaphore struct {
	semC chan struct{}
}

func NewSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}

func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.semC
}
