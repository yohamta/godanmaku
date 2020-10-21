package shot

import (
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

var (
	controllers = map[string]controller{
		"colorful": &colorful{baseController{}},
		"blue":     &blue{baseController{}},
		"laser":    &laser{baseController{}},
		"sparkle":  &sparkle{baseController{}},
	}
)

func PlayerShot(shooter Shooter, x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["blue"], shooter, x, y, degree)
}

func BlueLaser(shooter Shooter, x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["laser"], shooter, x, y, degree)
	s.spr = sprite.BlueLaser
}

func BlueLaserLong(shooter Shooter, x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["laser"], shooter, x, y, degree)
	s.spr = sprite.BlueLaserLong
}

func EnemyShot(shooter Shooter, x, y float64, degree int) {
	s := (*Shot)(shared.EnemyShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["colorful"], shooter, x, y, degree)
}
