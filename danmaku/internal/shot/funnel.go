package shot

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type funnel struct {
	baseController
}

func (c *funnel) init(s *Shot) {
	s.spr = sprite.PlayerBullet
	s.setSize(10, 10)
	s.collisionBox = collision.GetCollisionBox("NULL")
}

func (c *funnel) update(s *Shot) {
	s.sprIndex++
	if s.sprIndex >= s.spr.Length() {
		s.sprIndex = 0
	}
}

func (c *funnel) draw(s *Shot, screen *ebiten.Image) {
	spr := s.spr
	spr.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)
	spr.SetIndex(s.sprIndex)
	spr.Draw(screen)
}
