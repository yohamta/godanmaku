package shot

import (
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type playerNormalController struct {
	baseController
}

func (c *playerNormalController) init(s *Shot) {
	s.spr = sprite.PlayerBullet
	s.setSize(4, 4)
	s.setSpeed(2.56, s.degree)
}
