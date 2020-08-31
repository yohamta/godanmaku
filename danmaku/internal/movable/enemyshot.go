package movable

import (
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
)

const (
	enemyShotSize = 10
)

// EnemyShot represents player's bullet
type EnemyShot struct {
	ShObject
	shotSprite *sprite.Sprite
	isActive   bool
}

// NewEnemyShot returns initialized struct
func NewEnemyShot() *EnemyShot {
	base := &ShObject{}
	e := &EnemyShot{ShObject: *base}
	e.isActive = false

	return e
}

// IsActive returns if the actor is active in bool
func (e *EnemyShot) IsActive() bool {
	return e.isActive
}

// SetInactive returns if the actor is active in bool
func (e *EnemyShot) SetInactive() {
	e.isActive = false
}

// Init inits this
func (e *EnemyShot) Init(degree int, speed float64, x, y int) {
	e.speed = speed
	e.deg = degree

	e.vx = math.Cos(util.DegToRad(degree)) * speed
	e.vy = math.Sin(util.DegToRad(degree)) * speed

	e.width = enemyShotSize
	e.height = enemyShotSize
	e.x = float64(x)
	e.y = float64(y)

	e.shotSprite = sprite.RandomEnemyShot()

	e.isActive = true
}

// Draw draws this
func (e *EnemyShot) Draw(screen *ebiten.Image) {
	e.shotSprite.SetPosition(e.x, e.y)
	e.shotSprite.Draw(screen)
}

// Move moves this
func (e *EnemyShot) Move() {
	e.x = e.x + e.vx
	e.y = e.y + e.vy
	if e.isOutOfBoundary() {
		e.isActive = false
	}
}
