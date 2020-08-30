package actors

import (
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

// PlayerShot represents player's bullet
type PlayerShot struct {
	Actor
	spriteIndex int
	isActive    bool
}

// NewPlayerShot returns initialized struct
func NewPlayerShot() *PlayerShot {
	actor := &Actor{}
	p := &PlayerShot{Actor: *actor}
	p.isActive = false

	return p
}

// IsActive returns if the actor is active in bool
func (p *PlayerShot) IsActive() bool {
	return p.isActive
}

// SetInactive returns if the actor is active in bool
func (p *PlayerShot) SetInactive() {
	p.isActive = false
}

// Init inits this
func (p *PlayerShot) Init(degree int, speed float64, x, y, size int) {
	p.speed = speed
	p.deg = degree

	p.vx = math.Cos(degToRad(degree)) * speed
	p.vy = math.Sin(degToRad(degree)) * speed
	p.spriteIndex = degreeToDirectionIndex(degree)

	p.width = size
	p.height = size
	p.x = float64(x)
	p.y = float64(y)

	p.isActive = true
}

// Draw draws this
func (p *PlayerShot) Draw(screen *ebiten.Image) {
	sprite.PlayerBullet.SetPosition(p.x, p.y)
	sprite.PlayerBullet.SetIndex(p.spriteIndex)
	sprite.PlayerBullet.Draw(screen)
}

// Move moves this
func (p *PlayerShot) Move() {
	p.x = p.x + p.vx
	p.y = p.y + p.vy
	if p.isOutOfBoundary() {
		p.isActive = false
	}
}
