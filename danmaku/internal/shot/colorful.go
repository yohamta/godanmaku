package shot

import (
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type colorful struct {
	baseController
}

func (c *colorful) init(s *Shot) {
	s.spr = sprite.RandomEnemyShot()
	s.setSize(10, 10)
	s.setSpeed(1.44, s.degree)
}
