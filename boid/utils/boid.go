package utils

import (
	"boid/constants"
	"math/rand"
	"time"
)

type Boid struct {
	Pos      Vector
	velocity Vector
	id       int
}

var (
	Boids [constants.BoidCount]*Boid
)

// move the boid as per its velocity = [-1, 1]
func (b *Boid) moveOne() {
	// If boid reaches edge of board, bounce back by reversing the velocity direction
	next := b.Pos.Add(b.velocity)
	b.Pos = next
	if next.X >= constants.ScreenWidth || next.X < 0 {
		b.velocity = Vector{-b.velocity.X, b.velocity.Y}
	}
	if next.Y >= constants.ScreenHeight || next.Y < 0 {
		b.velocity = Vector{b.velocity.X, -b.velocity.Y}
	}
}

func (b *Boid) start() {
	// infinite loop as the boids should keep moving until the program terminates.
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func CreateBoid(idx int) {
	b := Boid{
		Pos: Vector{rand.Float64() * constants.ScreenWidth, rand.Float64() * constants.ScreenHeight},
		// Between -1 and 1 so that brids jump only 1 pixel at a time
		velocity: Vector{rand.Float64()*2 - 1.0, rand.Float64()*2 - 1.0},
		id:       idx,
	}
	Boids[idx] = &b
	// Start a green thread using goroutine
	go b.start()
}
