package actors

import (
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

var (
	sharedSprite *sprite.Sprite = nil
)

// PlayerBullet represents player of the game
type PlayerBullet struct {
	Actor
	vx          float64
	vy          float64
	spriteIndex int
}

// NewPlayerBullet returns initialized PlayerBullet
func NewPlayerBullet() *PlayerBullet {
	actor := &Actor{}
	p := &PlayerBullet{Actor: *actor}

	p.SetPosition(120, 160)
	p.SetSpeed(initPlayerSpeed)

	return p
}

// InitBullet inits this bullet
func (p *PlayerBullet) InitBullet(degree int, speed float64) {
	p.spriteIndex = DegreeToDirectionIndex(degree)
	p.Speed = speed
	// 	p.Vx = math.

	// nSpdX = ((Math.cos(Math.toRadians(degree))) * spd);
	// nSpdY = ((Math.sin(Math.toRadians(degree))) * spd);
}

// Draw draws this sprite
func (p *PlayerBullet) Draw(screen *ebiten.Image) {
	sprite.PlayerBullet.SetPosition(p.X, p.Y)
	sprite.PlayerBullet.SetIndex(p.spriteIndex)
	sprite.PlayerBullet.Draw(screen)
}

// Move moves player
func (p *PlayerBullet) Move(horizontal float64, vertical float64, isFire bool) {
	if vertical != 0 {
		p.vy = vertical * p.Speed
		p.Y = p.Y + p.vy
	}

	if horizontal != 0 {
		p.vx = horizontal * p.Speed
		p.X = p.X + p.vx
	}

	if vertical != 0 || horizontal != 0 {
		degree := RadToDeg(math.Atan2(vertical, horizontal))
		if isFire == false {
			p.SetDeg(degree)
		}
	}
}
