package input

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Input represents the state of user's input
type Input struct {
	Horizontal   float64
	Vertical     float64
	Fire         bool
	prevTickTime time.Time
}

// New creates new Input
func New() *Input {
	gameInput := &Input{}
	gameInput.prevTickTime = time.Now()
	return gameInput
}

// Update updates the input state
func (i *Input) Update() {
	i.takeKeyboardInput()
}

func (i *Input) takeKeyboardInput() {
	// Adjust sensitivity for keyboard input
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
