package shot

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type controller interface {
	init(s *Shot)
	update(s *Shot)
	draw(s *Shot, screen *ebiten.Image)
}

type baseController struct{}

func (c *baseController) init(s *Shot) {}

func (c *baseController) update(s *Shot) {}

func (c *baseController) draw(s *Shot, screen *ebiten.Image) {}

func (c *baseController) drawGrowEffect(s *Shot, screen *ebiten.Image) {
	sprite.Nova.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)

	w, _ := sprite.Nova.Size()
	scale := s.GetWidth() / float64(w) * 3

	sprite.Nova.DrawAdditive(screen, 0.5, scale, scale)
}
