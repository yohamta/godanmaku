package actors

import (
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

const (
	initPlayerSpeed = 2
	initPositionX   = 120
	initPositionY   = 160
)

// Player represents player of the game
type Player struct {
	Actor
}

// NewPlayer returns initialized Player
func NewPlayer() *Player {
	actor := &Actor{}
	p := &Player{Actor: *actor}

	p.SetPosition(initPositionX, initPositionY)
	p.SetSpeed(initPlayerSpeed)

	return p
}

// Draw draws this sprite
func (p *Player) Draw(screen *ebiten.Image) {
	sprite.Player.SetPosition(p.X, p.Y)
	sprite.Player.SetIndex(DegreeToDirectionIndex(p.Deg))
	sprite.Player.Draw(screen)
}

// Move moves player
func (p *Player) Move(horizontal float64, vertical float64, isFire bool) {
	if vertical != 0 {
		p.Vy = vertical * p.Speed
		p.Y = p.Y + p.Vy
	}

	if horizontal != 0 {
		p.Vx = horizontal * p.Speed
		p.X = p.X + p.Vx
	}

	if vertical != 0 || horizontal != 0 {
		if isFire == false {
			p.Deg = RadToDeg(math.Atan2(vertical, horizontal))
		}
	}
}
