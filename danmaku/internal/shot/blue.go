package shot

import (
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type blue struct {
	baseController
}

func (c *blue) init(s *Shot) {
	s.spr = sprite.PlayerBullet
	s.setSize(4, 4)
	s.setSpeed(2.56, s.degree)
}
