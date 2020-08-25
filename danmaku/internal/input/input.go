package input

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// GameInput represents the state of user's input
type GameInput struct {
	Horizontal   float64
	Vertical     float64
	Fire         bool
	prevTickTime time.Time
}

// New creates new GameInput
func New() *GameInput {
	gameInput := &GameInput{}
	gameInput.prevTickTime = time.Now()
	return gameInput
}

// Update updates the input state
func (i *GameInput) Update() {
	if time.Since(i.prevTickTime).Milliseconds() < 50 {
		return
	}
	i.prevTickTime = time.Now()

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		i.Vertical = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		i.Vertical = -1
	} else {
		i.Vertical = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.Horizontal = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.Horizontal = -1
	} else {
		i.Horizontal = 0
	}

	i.Fire = ebiten.IsKeyPressed(ebiten.KeySpace)
}
