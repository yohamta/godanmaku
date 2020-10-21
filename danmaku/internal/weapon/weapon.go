package weapon

import (
	"time"

	"github.com/yohamta/godanmaku/danmaku/internal/shot"
)

type Weapon interface {
	Fire(shooter shot.Shooter, x, y float64, degree int)
}

type ShotFactory func(shooter shot.Shooter, x, y float64, degree int)

type baseWeapon struct {
	shotFactory  ShotFactory
	lastShotTime time.Time
	playSound    bool
}
