package shot

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type colorful struct {
	baseController
}

func (c *colorful) init(s *Shot) {
	s.spr = sprite.RandomEnemyShot()
	s.setSize(10, 10)
	s.setSpeed(1.44, s.degree)
	s.collisionBox = collision.GetCollisionBox("WEAPON_NORMAL_1")
}

func (c *colorful) draw(s *Shot, screen *ebiten.Image) {
	spr := s.spr
	spr.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)
	spr.SetIndex(s.sprIndex)
	spr.Draw(screen)
}
