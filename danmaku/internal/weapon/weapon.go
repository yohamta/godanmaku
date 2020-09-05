package weapon

import (
	"github.com/yohamta/godanmaku/danmaku/internal/flyweight"
)

// Shooter represents shooter
type Shooter interface {
	GetX() float64
	GetY() float64
	GetDegree() int
}

// Weapon represents weapon
type Weapon interface {
	Fire(shooter Shooter, shots *flyweight.Pool)
}
