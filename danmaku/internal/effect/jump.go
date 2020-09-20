package effect

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type jump struct{ *baseController }

func (c *jump) init(e *Effect) {
}

func (c *jump) draw(e *Effect, screen *ebiten.Image) {
	if e.spriteFrame >= sprite.Jump.Length() {
		return
	}
	if e.updateCount < e.waitFrame {
		return
	}
	sprite.Jump.SetIndex(e.spriteFrame)
	sprite.Jump.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Jump.Draw(screen)

	scale := float64(sprite.Jump.GetWidth()) * e.scale *
		math.Min((1.-(float64(e.spriteFrame)/float64(sprite.Jump.Length()))+0.5), 1.)
	c.drawGrowEffect(e, scale, scale, 0.5, screen)
}

func (c *jump) update(e *Effect) {
	if e.updateCount < e.waitFrame {
		return
	}
	if e.updateCount == e.waitFrame {
		sound.PlaySe(sound.SeKindJump)
	}
	if e.updateCount > 0 && e.updateCount%5 == 0 {
		e.spriteFrame++
	}
	if e.spriteFrame >= sprite.Jump.Length() {
		e.isActive = false
		if e.callback != nil {
			e.callback()
			e.callback = nil
		}
	}
}
