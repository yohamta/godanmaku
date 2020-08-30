package shooting

import (
	"time"
)

const (
	weapon1ReloadTimeMs = 350
	weapon1Speed        = 8
	weapon1Size         = 4
)

// PlayerWeapon represents weapon
type PlayerWeapon interface {
	shot()
}

// PlayerWeapon1 represents player's weapon
type PlayerWeapon1 struct {
	lastShotTime time.Time
}

func (w *PlayerWeapon1) shot() {
	if time.Since(w.lastShotTime).Milliseconds() < weapon1ReloadTimeMs {
		return
	}
	w.lastShotTime = time.Now()
	x, y := player.getCenter()

	for i := 0; i < len(playerShots); i++ {
		s := playerShots[i]
		if s.isActive {
			continue
		}
		s.initBullet(player.deg, weapon1Speed, x, y, weapon1Size)
		break
	}

}
