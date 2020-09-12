package effect

import "github.com/hajimehoshi/ebiten"

// Effect represents the base of player, enemy, shots
type Effect struct {
	x           float64
	y           float64
	isActive    bool
	controller  controller
	updateCount int
	spriteFrame int
	waitFrame   int
	callback    func()
}

// NewEffect creates new effect
func NewEffect() *Effect {
	e := &Effect{}

	return e
}

// IsActive returns if this is active
func (e *Effect) IsActive() bool {
	return e.isActive
}

// Draw draws the player
func (e *Effect) Draw(screen *ebiten.Image) {
	e.controller.draw(e, screen)
}

// Update updates the effect
func (e *Effect) Update() {
	e.controller.update(e)
	e.updateCount++
}

func (e *Effect) init(c controller, x, y float64) {
	e.x = x
	e.y = y
	e.isActive = true
	e.controller = c
	e.updateCount = 0
	e.spriteFrame = 0
	e.waitFrame = 0
	c.init(e)
}
