package weapon

import (
	"time"

	"github.com/miyahoyo/godanmaku/danmaku/internal/shot"
)

type normal struct{ baseWeapon }

func (w *normal) Fire(shooter shot.Shooter, x, y float64, degree int) bool {
	if time.Since(w.lastShotTime).Milliseconds() < 350 {
		return false
	}
	w.lastShotTime = time.Now()
	w.shotFactory(shooter, x, y, degree)
	return true
}
