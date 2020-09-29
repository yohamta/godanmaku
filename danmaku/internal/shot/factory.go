package shot

import (
	"github.com/yotahamada/godanmaku/danmaku/internal/shared"
	"github.com/yotahamada/godanmaku/danmaku/internal/sprite"
)

var (
	controllers = map[string]controller{
		"colorful": &colorful{baseController{}},
		"blue":     &blue{baseController{}},
		"laser":    &laser{baseController{}},
		"sparkle":  &sparkle{baseController{}},
	}
)

// PlayerShot creates shot
func PlayerShot(x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["blue"], x, y, degree)
}

// BlueLaser creates shot
func BlueLaser(x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["laser"], x, y, degree)
	s.spr = sprite.BlueLaser
}

// BlueLaserLong creates shot
func BlueLaserLong(x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["laser"], x, y, degree)
	s.spr = sprite.BlueLaserLong
}

// EnemyShot creates shot
func EnemyShot(x, y float64, degree int) {
	s := (*Shot)(shared.EnemyShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["colorful"], x, y, degree)
}
