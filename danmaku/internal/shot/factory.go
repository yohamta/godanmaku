package shot

import "github.com/yohamta/godanmaku/danmaku/internal/shared"

var (
	playerNormal = &playerNormalController{baseController{}}
	enemyNormal  = &enemyNormalController{baseController{}}
)

// NormalPlayerShot creates shot
func NormalPlayerShot(x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(playerNormal, x, y, degree)
}

// NormalEnemyShot creates shot
func NormalEnemyShot(x, y float64, degree int) {
	s := (*Shot)(shared.EnemyShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(enemyNormal, x, y, degree)
}
