package effects

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

// Hit represents player of the game
type Hit struct {
	Effect
	lifeSpan int
}

// NewHit returns initialized Hit
func NewHit() *Hit {
	effect := &Effect{}
	h := &Hit{Effect: *effect}
	h.isActive = false

	return h
}

// StartEffect starts the effect
func (h *Hit) StartEffect(x, y float64) {
	h.lifeSpan = 20
	h.isActive = true
	h.x = float64(x)
	h.y = float64(y)
}

// Draw draws the player
func (h *Hit) Draw(screen *ebiten.Image) {
	sprite.Hit.SetPosition(h.x, h.y)
	sprite.Hit.Draw(screen)
}

// Update updates the effect
func (h *Hit) Update() {
	h.y--
	h.lifeSpan--
	if h.lifeSpan < 0 {
		h.isActive = false
	}
}
