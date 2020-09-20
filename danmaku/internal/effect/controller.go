package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type controller interface {
	init(e *Effect)
	update(e *Effect)
	draw(e *Effect, screen *ebiten.Image)
}

type baseController struct{}

func newBaseControlelr() *baseController {
	c := &baseController{}
	return c
}

func (c *baseController) init(e *Effect) {}

func (c *baseController) draw(e *Effect, screen *ebiten.Image) {}

func (c *baseController) update(e *Effect) {}

func (c *baseController) drawGrowEffect(e *Effect, width, height, strength float64, screen *ebiten.Image) {
	sprite.Nova.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)

	w, h := sprite.Nova.Size()
	scaleW := width / float64(w) * 2
	scaleH := height / float64(h) * 2

	sprite.Nova.DrawAdditive(screen, strength, scaleW, scaleH)
}
