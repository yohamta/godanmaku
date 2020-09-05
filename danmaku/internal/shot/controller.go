package shot

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
)

type controller interface {
	init(s *Shot)
	update(s *Shot)
	draw(s *Shot, screen *ebiten.Image)
}

type baseController struct{}

func (c *baseController) init(s *Shot) {
	// must be overriden
}

func (c *baseController) update(s *Shot) {
	s.setPosition(s.x+s.vx, s.y+s.vy)
	if util.IsOutOfArea(s, s.field) {
		s.isActive = false
	}
}

func (c *baseController) draw(s *Shot, screen *ebiten.Image) {
	spr := s.spr
	spr.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)
	spr.SetIndex(s.sprIndex)
	spr.Draw(screen)
}
