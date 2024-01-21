package main

import (
	"boid/constants"
	"boid/utils"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

// Render the state
func (g *Game) Draw(screen *ebiten.Image) {
	// print a diamond for each boid
	for _, boid := range utils.Boids {
		screen.Set(int(boid.Pos.X+1), int(boid.Pos.Y), green)
		screen.Set(int(boid.Pos.X-1), int(boid.Pos.Y), green)
		screen.Set(int(boid.Pos.X), int(boid.Pos.Y+1), green)
		screen.Set(int(boid.Pos.X), int(boid.Pos.Y-1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func main() {
	for i := 0; i < constants.BoidCount; i++ {
		utils.CreateBoid(i)
	}
	ebiten.SetWindowSize(constants.ScreenWidth*2, constants.ScreenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
