package utils

import (
	"boid/constants"
	"math"
	"math/rand"
	"sync"
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
	// There are 4 possible access in shared state (Read/Update map, update any boid's velocity/Pos). Add RWMutex for thread synchronization.
	rWlock = sync.RWMutex{}
)

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < constants.ViewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-constants.ViewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) calcAcceleration() Vector {
	upper, lower := b.Pos.AddNum(constants.ViewRadius), b.Pos.AddNum(-constants.ViewRadius)
	avgVelocity, avgPos, separation := Vector{0, 0}, Vector{0, 0}, Vector{0, 0}
	cnt := 0.0

	// Only reader's lock as we are not updating the shared state. We are only reading the map and other boid's velocity/Pos
	rWlock.RLock()
	for i := math.Max(0, lower.X); i <= math.Min(constants.ScreenWidth, upper.X); i++ {
		for j := math.Max(0, lower.Y); j <= math.Min(constants.ScreenHeight, upper.Y); j++ {
			if neighborBoid := BoidMap[int(i)][int(j)]; neighborBoid != -1 && neighborBoid != b.id {
				// all neighbors with current boid as centre of circle
				if dist := Boids[neighborBoid].Pos.Dist(b.Pos); dist < constants.ViewRadius {
					cnt++
					avgVelocity = avgVelocity.Add(Boids[neighborBoid].velocity)
					avgPos = avgPos.Add(Boids[neighborBoid].Pos)
					separation = separation.Add(b.Pos.Sub(Boids[neighborBoid].Pos).DivNum(dist))
				}
			}
		}
	}
	rWlock.RUnlock()

	accln := Vector{b.borderBounce(b.Pos.X, constants.ScreenWidth), b.borderBounce(b.Pos.Y, constants.ScreenHeight)}
	if cnt > 0 {
		avgVelocity, avgPos = avgVelocity.DivNum(cnt), avgPos.DivNum(cnt)
		acclnAlignment := avgVelocity.Sub(b.velocity).MulNum(constants.AdjRate)
		acclnCohesion := avgPos.Sub(b.Pos).MulNum(constants.AdjRate)
		acclnSeparation := separation.MulNum(constants.AdjRate)
		// Total accln based on alignment, cohesion and separation
		accln = accln.Add(acclnAlignment).Add(acclnCohesion).Add(acclnSeparation)
	}
	return accln
}

// move the boid as per its velocity = [-1, 1]
func (b *Boid) moveOne() {
	// Note: Moving this out of Add() on line 62 otherwise thread will be blocked as it can't acquire same lock twice.
	accln := b.calcAcceleration()

	// Writer's lock as we are updating the shared state. We are updating the current boid's velocity/Pos and the map
	rWlock.Lock()
	// converge boid with neighboring boids
	b.velocity = b.velocity.Add(accln).Limit(-1, 1)
	// reset current pos to -1
	BoidMap[int(b.Pos.X)][int(b.Pos.Y)] = -1
	// If boid reaches edge of board, bounce back by reversing the velocity direction
	next := b.Pos.Add(b.velocity)
	b.Pos = next
	BoidMap[int(b.Pos.X)][int(b.Pos.Y)] = b.id
	rWlock.Unlock()
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
