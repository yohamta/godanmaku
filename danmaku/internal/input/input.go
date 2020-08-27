package input

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	joyStickRadius float64 = 50
)

type point struct{ x, y int }

// Input represents the state of user's input
type Input struct {
	Horizontal      float64
	Vertical        float64
	Fire            bool
	prevTickTime    time.Time
	joyStickTouchID int
	joyStickCenter  point
	isTouch         bool
}

// New creates new Input
func New() *Input {
	gameInput := &Input{}
	gameInput.prevTickTime = time.Now()
	gameInput.joyStickTouchID = -1
	return gameInput
}

// Update updates the input state
func (input *Input) Update() {
	input.readTouchInput()
	input.readKeyboardInput()
}

// IsJoyStickUsing returns if joystick is now using
func (input *Input) IsJoyStickUsing() bool {
	return input.joyStickTouchID != -1
}

// GetJoyStickPosition returns JoyStickPosition
func (input *Input) GetJoyStickPosition() (int, int) {
	return input.joyStickCenter.x, input.joyStickCenter.y
}

func (input *Input) readTouchInput() {
	touchIds := ebiten.TouchIDs()
	justPressedTouchIds := inpututil.JustPressedTouchIDs()
	if input.joyStickTouchID == -1 && justPressedTouchIds != nil {
		// started using joystick
		input.joyStickTouchID = justPressedTouchIds[0]
		x, y := ebiten.TouchPosition(input.joyStickTouchID)
		input.joyStickCenter = point{x: x, y: y}
	}

	for i := 0; i < len(touchIds); i++ {
		tid := touchIds[i]
		if input.joyStickTouchID == tid {
			// joystick control
			x, y := ebiten.TouchPosition(tid)
			dx := x - input.joyStickCenter.x
			dy := y - input.joyStickCenter.y
			input.Horizontal = float64(dx) / joyStickRadius
			input.Vertical = float64(dy) / joyStickRadius
		}
	}

	if inpututil.IsTouchJustReleased(input.joyStickTouchID) {
		// Released joystick
		input.joyStickTouchID = -1
		input.Horizontal = 0
		input.Vertical = 0
	}
}

func (input *Input) readKeyboardInput() {
	// Adjust sensitivity for keyboard input
	if time.Since(input.prevTickTime).Milliseconds() < 50 {
		return
	}
	input.prevTickTime = time.Now()

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		input.Vertical = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		input.Vertical = -1
	} else {
		input.Vertical = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		input.Horizontal = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		input.Horizontal = -1
	} else {
		input.Horizontal = 0
	}

	input.Fire = ebiten.IsKeyPressed(ebiten.KeySpace)
}
