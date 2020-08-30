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
	playerWidth     = 10
	playerHeight    = 10
)

// Player represents player of the game
type Player struct {
	Actor
}

// NewPlayer returns initialized Player
func NewPlayer() *Player {
	actor := &Actor{}
	p := &Player{Actor: *actor}

	p.width = playerWidth
	p.height = playerHeight
	p.setPosition(initPositionX, initPositionY)
	p.setSpeed(initPlayerSpeed)

	return p
}

// Draw draws the player
func (p *Player) Draw(screen *ebiten.Image) {
	sprite.Player.SetPosition(p.x, p.y)
	sprite.Player.SetIndex(degreeToDirectionIndex(p.deg))
	sprite.Player.Draw(screen)
}

// Move moves the player
func (p *Player) Move(horizontal float64, vertical float64, isFire bool) {
	if vertical != 0 {
		p.vy = vertical * p.speed
		p.y = p.y + p.vy
		p.y = math.Max(float64(boundarizer.GetTop()+p.height/2), p.y)
		p.y = math.Min(float64(boundarizer.GetBottom()-p.height/2), p.y)
	}

	if horizontal != 0 {
		p.vx = horizontal * p.speed
		p.x = p.x + p.vx
		p.x = math.Max(float64(boundarizer.GetLeft()+p.width/2), p.x)
		p.x = math.Min(float64(boundarizer.GetRight()-p.width/2), p.x)
	}

	if vertical != 0 || horizontal != 0 {
		if isFire == false {
			p.deg = radToDeg(math.Atan2(vertical, horizontal))
		}
	}

}
