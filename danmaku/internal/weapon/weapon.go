package weapon

import (
	"time"
)

type Weapon interface {
	Fire(x, y float64, degree int)
}

type shotFactory func(x, y float64, degree int)

type baseWeapon struct {
	shotFactory  shotFactory
	lastShotTime time.Time
	playSound    bool
}
