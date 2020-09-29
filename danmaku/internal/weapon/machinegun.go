package weapon

import (
	"time"

	"github.com/yotahamada/godanmaku/danmaku/internal/sound"
)

type machingun struct{ baseWeapon }

// Fire create shots
func (w *machingun) Fire(x, y float64, degree int) {
	if time.Since(w.lastShotTime).Milliseconds() < 75 {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, degree)
	if w.playSound {
		sound.PlaySe(sound.SeKindShot)
	}
}
