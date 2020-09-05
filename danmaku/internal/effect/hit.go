package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type hit struct{}

func (c *hit) init(e *Effect) {}

func (c *hit) draw(e *Effect, screen *ebiten.Image) {
	sprite.Hit.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Hit.Draw(screen)
}

func (c *hit) update(e *Effect) {
	e.y--
	if e.updateCount >= 20 {
		e.isActive = false
	}
}
