package weapon

import (
	"time"

	"github.com/yohamta/godanmaku/danmaku/internal/shot"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"
)

type normal struct{ baseWeapon }

func (w *normal) Fire(shooter shot.Shooter, x, y float64, degree int) {
	if time.Since(w.lastShotTime).Milliseconds() < 350 {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(shooter, x, y, degree)
	if w.playSound {
		sound.PlaySe(sound.SeKindShot)
	}
}
