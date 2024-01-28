package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxx = 5

var (
	lock [maxx]sync.Mutex
)

type Philosopher interface {
	Eat(threadName string)
	Think(threadName string)
}

type philosopher struct {
	id        int
	leftFork  *sync.Mutex
	rightFork *sync.Mutex
}

func NewPhilosopher(id int) Philosopher {
	return &philosopher{
		id:        id,
		leftFork:  &lock[(id)%maxx],
		rightFork: &lock[(id+1)%maxx],
	}
}

// Slow eating process
func (p *philosopher) Eat(threadName string) {
	// pickup the forks -- first left then right
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Printf("%v is eating\n", threadName)
	time.Sleep(5*time.Second + time.Duration(rand.Intn(5))*time.Second)
	fmt.Printf("%v finished eating\n", threadName)

	// put down the forks
	p.leftFork.Unlock()
	p.rightFork.Unlock()
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
