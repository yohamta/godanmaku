package shared

import (
	"github.com/yohamta/godanmaku/danmaku/internal/flyweight"
)

var (
	// PlayerShots is a pool
	PlayerShots *flyweight.Pool = flyweight.NewPool()
	// EnemyShots is a pool
	EnemyShots *flyweight.Pool = flyweight.NewPool()
	// Effects is a pool
	Effects *flyweight.Pool = flyweight.NewPool()
	// Enemies is a pool
	Enemies *flyweight.Pool = flyweight.NewPool()

	// OffsetX is screen offset
	OffsetX float64
	// OffsetY is screen offset
	OffsetY float64
)
