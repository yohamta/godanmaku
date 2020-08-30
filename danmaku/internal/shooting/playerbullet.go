package shooting

import (
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

var (
	sharedSprite *sprite.Sprite = nil
)

// PlayerBulletState represents player bullet state
type PlayerBulletState int

const (
	PlayerBulletStateActive PlayerBulletState = iota
	PlayerBulletStateInActive
)

// PlayerBullet represents player's bullet
type PlayerBullet struct {
	Actor
	spriteIndex int
	state       PlayerBulletState
}

// NewPlayerBullet returns initialized struct
func NewPlayerBullet() *PlayerBullet {
	actor := &Actor{}
	p := &PlayerBullet{Actor: *actor}
	p.state = PlayerBulletStateInActive

	p.setPosition(120, 160)
	p.setSpeed(initPlayerSpeed)

	return p
}

// InitBullet inits this bullet
func (p *PlayerBullet) initBullet(degree int, speed float64) {
	p.speed = speed
	p.vx = math.Cos(degToRad(degree)) * speed
	p.vy = math.Sin(degToRad(degree)) * speed
	p.spriteIndex = degreeToDirectionIndex(degree)
	p.state = PlayerBulletStateActive
}

// Draw draws this sprite
func (p *PlayerBullet) draw(screen *ebiten.Image) {
	sprite.PlayerBullet.SetPosition(p.x, p.y)
	sprite.PlayerBullet.SetIndex(p.spriteIndex)
	sprite.PlayerBullet.Draw(screen)
}

// Move moves this
func (p *PlayerBullet) Move() {
	p.x = p.x + p.vx
	p.y = p.y + p.vy
}
