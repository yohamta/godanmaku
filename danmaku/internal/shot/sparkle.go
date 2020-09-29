package shot

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/godanmaku/danmaku/internal/collision"
	"github.com/yotahamada/godanmaku/danmaku/internal/shared"
	"github.com/yotahamada/godanmaku/danmaku/internal/sprite"
)

type sparkle struct {
	baseController
}

func (c *sparkle) init(s *Shot) {
	s.spr = sprite.Sparkle
	s.setSize(10, 10)
	s.setSpeed(1.44, s.degree)
	s.sprIndex = rand.Intn(54)
	s.collisionBox = collision.GetCollisionBox("WEAPON_NORMAL_1")
}

func (c *sparkle) update(s *Shot) {
	s.sprIndex++
	if s.sprIndex >= 55 {
		s.sprIndex = 0
	}
}

func (c *sparkle) draw(s *Shot, screen *ebiten.Image) {
	spr := s.spr
	spr.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)
	spr.SetIndex(s.sprIndex)
	spr.Draw(screen)
}
