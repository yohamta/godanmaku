package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type locus struct{}

func (c *locus) init(e *Effect) {
}

func (c *locus) draw(e *Effect, screen *ebiten.Image) {
	if e.updateCount < e.waitFrame {
		return
	}
	if e.spriteFrame >= sprite.Locus.Length() {
		return
	}
	sprite.Locus.SetIndex(e.spriteFrame)
	sprite.Locus.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Locus.DrawWithScale(screen, 1)
}

func (c *locus) update(e *Effect) {
	if e.updateCount < e.waitFrame {
		return
	}
	if e.updateCount > 0 && e.updateCount%5 == 0 {
		e.spriteFrame++
	}
	if e.spriteFrame >= sprite.Locus.Length() {
		e.isActive = false
	}
}
