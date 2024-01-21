package utils

import (
	"boid/constants"
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	Pos      Vector
	velocity Vector
	id       int
}

var (
	Boids   [constants.BoidCount]*Boid
	BoidMap [constants.ScreenWidth + 1][constants.ScreenHeight + 1]int
)

func (b *Boid) calcAcceleration() Vector {
	upper, lower := b.Pos.AddNum(constants.ViewRadius), b.Pos.AddNum(-constants.ViewRadius)
	avgVelocity := Vector{0, 0}
	cnt := 0.0
	for i := math.Max(0, lower.X); i <= math.Min(constants.ScreenWidth, upper.X); i++ {
		for j := math.Max(0, lower.Y); j <= math.Min(constants.ScreenHeight, upper.Y); j++ {
			if neighborBoid := BoidMap[int(i)][int(j)]; neighborBoid != -1 && neighborBoid != b.id {
				// all neighbors with current boid as centre of circle
				if dist := Boids[neighborBoid].Pos.Dist(b.Pos); dist < constants.ViewRadius {
					cnt++
					avgVelocity = avgVelocity.Add(Boids[neighborBoid].velocity)
				}
			}
		}
	}
	accln := Vector{0, 0}
	if cnt > 0 {
		avgVelocity = avgVelocity.DivNum(cnt)
		accln = avgVelocity.Sub(b.velocity).MulNum(constants.AdjRate)
	}
	return accln
}

// move the boid as per its velocity = [-1, 1]
func (b *Boid) moveOne() {
	// converge boid with neighboring boids
	b.velocity = b.velocity.Add(b.calcAcceleration()).Limit(-1, 1)
	// reset current pos to -1
	BoidMap[int(b.Pos.X)][int(b.Pos.Y)] = -1
	// If boid reaches edge of board, bounce back by reversing the velocity direction
	next := b.Pos.Add(b.velocity)
	b.Pos = next
	BoidMap[int(b.Pos.X)][int(b.Pos.Y)] = b.id

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
		// wait a bit for other goroutines to make progress. Otherwise the motion won't be smooth.
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
	BoidMap[int(b.Pos.X)][int(b.Pos.Y)] = b.id
	// Start a green thread using goroutine
	go b.start()
}
