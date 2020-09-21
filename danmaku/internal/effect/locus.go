package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type locus struct{ *baseController }

func (c *locus) init(e *Effect) {
	e.waitFrame = 3
}

func (c *locus) draw(e *Effect, screen *ebiten.Image) {
	if e.updateCount < e.waitFrame {
		return
	}
	if e.spriteFrame >= sprite.Backfire.Length() {
		return
	}
	sprite.Backfire.SetIndex(e.spriteFrame)
	sprite.Backfire.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Backfire.Draw(screen)
}

func (c *locus) update(e *Effect) {
	if e.updateCount < e.waitFrame {
		return
	}
	if e.updateCount%4 == 0 {
		e.spriteFrame++
	}
	if e.spriteFrame >= sprite.Backfire.Length() {
		e.isActive = false
	}
}
