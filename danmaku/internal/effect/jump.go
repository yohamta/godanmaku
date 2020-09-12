package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type jump struct{}

func (c *jump) init(e *Effect) {}

func (c *jump) draw(e *Effect, screen *ebiten.Image) {
	if e.spriteFrame >= sprite.Jump.Length() {
		return
	}
	sprite.Jump.SetIndex(e.spriteFrame)
	sprite.Jump.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Jump.Draw(screen)
}

func (c *jump) update(e *Effect) {
	if e.updateCount > 0 && e.updateCount%3 == 0 {
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
