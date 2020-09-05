package shot

import (
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type enemyNormalController struct {
	baseController
}

func (c *enemyNormalController) init(s *Shot) {
	s.spr = sprite.RandomEnemyShot()
	s.setSize(10, 10)
	s.setSpeed(1.44, s.degree)
}
