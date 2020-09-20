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
	jumpController      = new(jump)
	locusController     = new(locus)
)

// CreateLocusEffect creates an effect
func CreateLocusEffect(x, y float64) {
	e := (*Effect)(shared.BackEffects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(locusController, x, y)
	e.waitFrame = 3
}

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

// CreateJump creates an effect
func CreateJump(x, y float64, wait int, callback func()) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(jumpController, x, y)
	e.waitFrame = wait
	e.callback = callback
}
