package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
)

// Controller represents effect controller
type Controller interface {
	init(e *Effect)
	update(e *Effect)
	draw(e *Effect, screen *ebiten.Image)
}

var (
	hit       = new(HitController)
	explosion = new(ExplosionController)
)

// CreateHitEffect creates an effect
func CreateHitEffect(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(hit, x, y)
}

// CreateExplosion creates an effect
func CreateExplosion(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(explosion, x, y)
}
