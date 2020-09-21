package effect

import (
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

var (
	normalController = &normal{newBaseControlelr()}
)

// CreateLocusEffect creates an effect
func CreateLocusEffect(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(normalController, x, y)
	e.sprite = sprite.Backfire
	e.waitFrame = 3
	e.fps = 10
}

// CreateHitEffect creates an effect
func CreateHitEffect(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(normalController, x, y)
	e.sprite = sprite.Hit
	e.se = sound.SeKindHit2
	e.fps = 15
}

// CreateExplosion creates an effect
func CreateExplosion(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(normalController, x, y)
	e.sprite = sprite.Explosion
	e.se = sound.SeKindBomb
	e.fps = 20
}

// CreateJump creates an effect
func CreateJump(x, y float64, wait int, callback func()) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(normalController, x, y)
	e.sprite = sprite.Jump
	e.se = sound.SeKindJump
	e.waitFrame = wait
	e.callback = callback
	e.fps = 12
}
