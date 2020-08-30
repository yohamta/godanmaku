package weapons

import (
	"time"

	"github.com/yohamta/godanmaku/danmaku/internal/shooting/actors"
)

const (
	weapon1ReloadTimeMs = 350
	weapon1Speed        = 2.56
	weapon1Size         = 4
)

// PlayerWeapon1 represents player's weapon
type PlayerWeapon1 struct {
	lastShotTime time.Time
}

// Shot create shots
func (w *PlayerWeapon1) Shot(x, y float64, degree int, playerShots []*actors.PlayerBullet) {
	if time.Since(w.lastShotTime).Milliseconds() < weapon1ReloadTimeMs {
		return
	}
	w.lastShotTime = time.Now()

	for i := 0; i < len(playerShots); i++ {
		s := playerShots[i]
		if s.IsActive() {
			continue
		}
		s.InitPlayerShot(degree, weapon1Speed, int(x), int(y), weapon1Size)
		break
	}

}
