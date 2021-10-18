package shot

import (
	"github.com/miyahoyo/godanmaku/danmaku/internal/shared"
	"github.com/miyahoyo/godanmaku/danmaku/internal/sprite"
)

var (
	controllers = map[string]controller{
		"item":     &item{baseController{}},
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

func PowerUpItem(x, y float64, degree int) {
	s := (*Shot)(shared.Items.CreateFromPool())
	if s == nil {
		return
	}
	s.itemKind = ItemKindPowerUp
	s.init(controllers["item"], nil, x, y, degree)
}

func RecoveryItem(x, y float64, degree int) {
	s := (*Shot)(shared.Items.CreateFromPool())
	if s == nil {
		return
	}
	s.itemKind = ItemKindRecovery
	s.init(controllers["item"], nil, x, y, degree)
}
