package weapon

import (
	"time"

	"github.com/miyahoyo/godanmaku/danmaku/internal/shot"
)

type machingun struct{ baseWeapon }

func (w *machingun) Fire(shooter shot.Shooter, x, y float64, degree int) bool {
	if time.Since(w.lastShotTime).Milliseconds() < 75 {
		return false
	}
	w.lastShotTime = time.Now()
	w.shotFactory(shooter, x, y, degree)
	return true
}
