package input

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// GameInput represents the state of user's input
type GameInput struct {
	Up           float64
	Left         float64
	Down         float64
	Right        float64
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

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		i.Up = 1
	} else {
		i.Up = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.Left = 1
	} else {
		i.Left = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		i.Down = 1
	} else {
		i.Down = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.Right = 1
	} else {
		i.Right = 0
	}

	i.Fire = ebiten.IsKeyPressed(ebiten.KeySpace)
}
