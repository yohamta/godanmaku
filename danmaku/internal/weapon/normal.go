package weapon

import (
	"time"
)

type normal struct{ baseWeapon }

// Fire create shots
func (w *normal) Fire(x, y float64, degree int) {
	if time.Since(w.lastShotTime).Milliseconds() < 350 {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, degree)
}
