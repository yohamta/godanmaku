package shared

import (
	"image"

	"github.com/yohamta/godanmaku/danmaku/internal/flyweight"
)

type Input struct {
	Horizontal float64
	Vertical   float64
	Fire       bool
}

var (
	ScreenSize image.Point

	PlayerShots   *flyweight.Pool = flyweight.NewPool()
	EnemyShots    *flyweight.Pool = flyweight.NewPool()
	PlayerFunnels *flyweight.Pool = flyweight.NewPool()
	BackEffects   *flyweight.Pool = flyweight.NewPool()
	Effects       *flyweight.Pool = flyweight.NewPool()
	Enemies       *flyweight.Pool = flyweight.NewPool()

	OffsetX float64
	OffsetY float64

	GameInput Input
)
