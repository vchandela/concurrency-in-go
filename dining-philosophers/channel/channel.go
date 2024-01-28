package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxx = 5

var (
	fork = make([]chan struct{}, maxx)
)

type Philosopher interface {
	Eat(threadName string)
	Think(threadName string)
}

type philosopher struct {
	id        int
	leftFork  chan struct{}
	rightFork chan struct{}
}

func NewPhilosopher(id int) Philosopher {
	return &philosopher{
		id:        id,
		leftFork:  fork[(id)%maxx],
		rightFork: fork[(id+1)%maxx],
	}
}

// Slow eating process
func (p *philosopher) Eat(threadName string) {
	// pickup the forks -- first left then right
	p.leftFork <- struct{}{}
	p.rightFork <- struct{}{}

	fmt.Printf("%v is eating\n", threadName)
	time.Sleep(5*time.Second + time.Duration(rand.Intn(5))*time.Second)
	fmt.Printf("%v finished eating\n", threadName)

	// put down the forks
	<-p.leftFork
	<-p.rightFork
}

// Fast thinking process
func (p *philosopher) Think(threadName string) {
	fmt.Printf("%v is thinking\n", threadName)
	time.Sleep(2*time.Second + time.Duration(rand.Intn(2))*time.Second)
	fmt.Printf("%v finished thinking\n", threadName)

	// start eating
	p.Eat(threadName)
}

func meeting(i int) {
	threadName := fmt.Sprintf("philosopher %v", i)
	p := NewPhilosopher(i)
	// Each philosopher starts thinking
	p.Think(threadName)
}

func Run() {
	for i := 0; i < maxx; i++ {
		fork[i] = make(chan struct{}, 1)
	}
	var wg sync.WaitGroup
	for i := 0; i < maxx; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			meeting(i)
		}(&wg, i)
	}
	wg.Wait()
}
