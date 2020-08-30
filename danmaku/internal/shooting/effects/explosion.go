package effects

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

// Explosion represents player of the game
type Explosion struct {
	Effect
	spriteFrame int
	updateCount int
}

// NewExplosion returns initialized Explosion
func NewExplosion() *Explosion {
	effect := &Effect{}
	h := &Explosion{Effect: *effect}
	h.isActive = false

	return h
}

// StartEffect starts the effect
func (h *Explosion) StartEffect(x, y int) {
	h.spriteFrame = 0
	h.updateCount = 0
	h.isActive = true
	h.x = float64(x)
	h.y = float64(y)
}

// Draw draws the player
func (h *Explosion) Draw(screen *ebiten.Image) {
	if h.spriteFrame >= sprite.Explosion.Length() {
		return
	}
	sprite.Explosion.SetIndex(h.spriteFrame)
	sprite.Explosion.SetPosition(h.x, h.y)
	sprite.Explosion.Draw(screen)
}

// Update updates the effect
func (h *Explosion) Update() {
	h.updateCount++
	if h.updateCount > 0 && h.updateCount%2 == 0 {
		h.spriteFrame++
	}
	if h.spriteFrame >= sprite.Explosion.Length() {
		h.isActive = false
	}
}
