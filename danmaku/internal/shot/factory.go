package shot

import "github.com/yohamta/godanmaku/danmaku/internal/shared"

var (
	controllers = map[string]controller{
		"colorful":  &colorful{baseController{}},
		"blue":      &blue{baseController{}},
		"blueLaser": &blueLaser{baseController{}},
	}
)

// PlayerShot creates shot
func PlayerShot(x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["blue"], x, y, degree)
}

// BlueLaser creates shot
func BlueLaser(x, y float64, degree int) {
	s := (*Shot)(shared.PlayerShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["blueLaser"], x, y, degree)
}

// EnemyShot creates shot
func EnemyShot(x, y float64, degree int) {
	s := (*Shot)(shared.EnemyShots.CreateFromPool())
	if s == nil {
		return
	}
	s.init(controllers["colorful"], x, y, degree)
}
