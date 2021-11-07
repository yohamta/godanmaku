package shot

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type blue struct {
	baseController
}

func (c *blue) init(s *Shot) {
	s.spr = sprite.Get("PSHOT_1")
	s.setSize(10, 10)
	s.setSpeed(3.56, s.degree)
	s.collisionBox = collision.GetCollisionBox("WEAPON_NORMAL_1")
	s.sprIndex = 0
}

func (c *blue) draw(s *Shot, screen *ebiten.Image) {
	x, y := s.x-shared.OffsetX, s.y-shared.OffsetY
	ganim8.DrawSprite(screen, s.spr, s.sprIndex, x, y, 0, 1, 1, .5, .5)
}
