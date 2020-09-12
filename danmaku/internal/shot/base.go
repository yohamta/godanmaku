package shot

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
)

type controller interface {
	init(s *Shot)
	update(s *Shot)
	draw(s *Shot, screen *ebiten.Image)
}

type baseController struct{}

func (c *baseController) init(s *Shot) {}

func (c *baseController) update(s *Shot) {}

func (c *baseController) draw(s *Shot, screen *ebiten.Image) {
	spr := s.spr
	spr.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)
	spr.SetIndex(s.sprIndex)
	spr.Draw(screen)
}
