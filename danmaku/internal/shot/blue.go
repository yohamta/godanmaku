package shot

import (
	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type blue struct {
	baseController
}

func (c *blue) init(s *Shot) {
	s.spr = sprite.PlayerBullet
	s.setSize(10, 10)
	s.setSpeed(3.56, s.degree)
	s.collisionBox = collision.GetCollisionBox("WEAPON_NORMAL_1")
}
