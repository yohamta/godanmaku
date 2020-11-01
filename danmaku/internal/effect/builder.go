package effect

import (
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

var (
	normalController = &normal{newBaseControlelr()}
)

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

func CreateHitEffect(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(normalController, x, y)
	e.sprite = sprite.Hit
	e.fps = 15
}

func CreateHitLargeEffect(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(normalController, x, y)
	e.sprite = sprite.Hit
	e.scale = 2
	e.fps = 15
}

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

func CreateJump(x, y float64, wait int, callback func()) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(normalController, x, y)
	e.sprite = sprite.Jump
	e.waitFrame = wait
	e.callback = callback
	e.callbackFrame = 3
	e.fps = 12
}
