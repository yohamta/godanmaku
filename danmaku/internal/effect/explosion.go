package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type explosion struct{}

func (c *explosion) init(e *Effect) {}

func (c *explosion) draw(e *Effect, screen *ebiten.Image) {
	if e.spriteFrame >= sprite.Explosion.Length() {
		return
	}
	sprite.Explosion.SetIndex(e.spriteFrame)
	sprite.Explosion.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Explosion.Draw(screen)
}

func (c *explosion) update(e *Effect) {
	if e.updateCount > 0 && e.updateCount%2 == 0 {
		e.spriteFrame++
	}
	if e.spriteFrame >= sprite.Explosion.Length() {
		e.isActive = false
	}
}
