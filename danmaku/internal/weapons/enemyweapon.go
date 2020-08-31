package weapons

import (
	"math/rand"

	"github.com/yohamta/godanmaku/danmaku/internal/movable"
)

const (
	enemyShotSpeed = 1.44
)

// EnemyAttack make enemy attack the player
func EnemyAttack(enemy *movable.Enemy, player *movable.Player, enemyShots []*movable.EnemyShot) {

	for i := 0; i < len(enemyShots); i++ {
		s := enemyShots[i]
		if s.IsActive() {
			continue
		}

		blur := int(rand.Float64() * 20)
		if rand.Float64() < 0.5 {
			blur *= -1
		}
		s.Init(enemy.GetDeg()+blur, enemyShotSpeed, enemy.GetX(), enemy.GetY())
		break
	}

}
