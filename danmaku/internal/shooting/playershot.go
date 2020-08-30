package shooting

import (
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

var (
	sharedSprite *sprite.Sprite = nil
)

// PlayerBullet represents player's bullet
type PlayerBullet struct {
	Actor
	spriteIndex int
	isActive    bool
}

// NewPlayerShot returns initialized struct
func NewPlayerShot() *PlayerBullet {
	actor := &Actor{}
	p := &PlayerBullet{Actor: *actor}
	p.isActive = false

	p.setPosition(120, 160)
	p.setSpeed(initPlayerSpeed)

	return p
}

func (p *PlayerBullet) initBullet(degree int, speed float64, x, y, size int) {
	p.speed = speed
	p.vx = math.Cos(degToRad(degree)) * speed
	p.vy = math.Sin(degToRad(degree)) * speed
	p.spriteIndex = degreeToDirectionIndex(degree)

	p.width = size
	p.height = size
	p.x = float64(x)
	p.y = float64(y)

	p.isActive = true
}

func (p *PlayerBullet) draw(screen *ebiten.Image) {
	sprite.PlayerBullet.SetPosition(p.x, p.y)
	sprite.PlayerBullet.SetIndex(p.spriteIndex)
	sprite.PlayerBullet.Draw(screen)
}

func (p *PlayerBullet) move(boundary Boundary) {
	p.x = p.x + p.vx
	p.y = p.y + p.vy
	if p.isOutOfBoundary(boundary) {
		p.isActive = false
	}
}
