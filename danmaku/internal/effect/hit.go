package effect

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

// HitController represents player of the game
type HitController struct{}

func (c *HitController) init(e *Effect) {}

// Draw draws the player
func (c *HitController) draw(e *Effect, screen *ebiten.Image) {
	sprite.Hit.SetPosition(e.x, e.y)
	sprite.Hit.Draw(screen)
}

// Update updates the effect
func (c *HitController) update(e *Effect) {
	e.y--
	if e.updateCount >= 20 {
		e.isActive = false
	}
}
