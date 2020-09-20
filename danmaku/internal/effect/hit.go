package effect

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type hit struct{ *baseController }

func (c *hit) init(e *Effect) {
	e.scale = rand.Float64()*1.5 + 0.5
}

func (c *hit) draw(e *Effect, screen *ebiten.Image) {
	if e.spriteFrame >= sprite.Hit.Length() {
		return
	}
	sprite.Hit.SetIndex(e.spriteFrame)
	sprite.Hit.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Hit.DrawWithScale(screen, e.scale)
}

func (c *hit) update(e *Effect) {
	if e.updateCount == 0 {
		sound.PlaySe(sound.SeKindHit2)
	}
	if e.updateCount > 0 && e.updateCount%4 == 0 {
		e.spriteFrame++
	}
	if e.spriteFrame >= sprite.Hit.Length() {
		e.isActive = false
	}
}
