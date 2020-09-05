package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
)

type controller interface {
	init(e *Effect)
	update(e *Effect)
	draw(e *Effect, screen *ebiten.Image)
}

var (
	hitController       = new(hit)
	explosionController = new(explosion)
)

// CreateHitEffect creates an effect
func CreateHitEffect(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(hitController, x, y)
}

// CreateExplosion creates an effect
func CreateExplosion(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(explosionController, x, y)
}
