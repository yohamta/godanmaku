package shot

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
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
	x, y := s.x-shared.OffsetX, s.y-shared.OffsetY
	ganim8.DrawSprite(screen, s.spr, s.sprIndex, x, y, 0, 1, 1, .5, .5)
}
