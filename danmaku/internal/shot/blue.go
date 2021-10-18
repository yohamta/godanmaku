package shot

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miyahoyo/godanmaku/danmaku/internal/collision"
	"github.com/miyahoyo/godanmaku/danmaku/internal/shared"
	"github.com/miyahoyo/godanmaku/danmaku/internal/sprite"
)

type blue struct {
	baseController
}

func (c *blue) init(s *Shot) {
	s.spr = sprite.PlayerBullet
	s.setSize(10, 10)
	s.setSpeed(3.56, s.degree)
	s.collisionBox = collision.GetCollisionBox("WEAPON_NORMAL_1")
	s.sprIndex = 0
}

func (c *blue) draw(s *Shot, screen *ebiten.Image) {
	spr := s.spr
	spr.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)
	spr.SetIndex(s.sprIndex)
	spr.Draw(screen)
}
