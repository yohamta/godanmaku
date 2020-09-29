package shot

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/godanmaku/danmaku/internal/collision"
	"github.com/yotahamada/godanmaku/danmaku/internal/shared"
)

type laser struct {
	baseController
}

func (c *laser) init(s *Shot) {
	s.setSize(16, 16)
	s.setSpeed(3.56, s.degree)

	degIndex := int((s.degree + 3600) % 360 / 15)
	degBy15 := degIndex * 15
	s.collisionBox = collision.GetCollisionBox(laserCollisionIDMap[degBy15])
	s.sprIndex = degIndex

	adjust := laserAdjustMap[degBy15]
	s.setPosition(s.x+adjust.x, s.y+adjust.y)
}

func (c *laser) draw(s *Shot, screen *ebiten.Image) {
	spr := s.spr
	spr.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)
	spr.SetIndex(s.sprIndex)
	spr.Draw(screen)
}
