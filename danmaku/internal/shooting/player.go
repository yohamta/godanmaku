package shooting

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

	p.setPosition(initPositionX, initPositionY)
	p.setSpeed(initPlayerSpeed)

	return p
}

func (p *Player) draw(screen *ebiten.Image) {
	sprite.Player.SetPosition(p.x, p.y)
	sprite.Player.SetIndex(degreeToDirectionIndex(p.deg))
	sprite.Player.Draw(screen)
}

func (p *Player) move(horizontal float64, vertical float64, isFire bool) {
	if vertical != 0 {
		p.vy = vertical * p.speed
		p.y = p.y + p.vy
	}

	if horizontal != 0 {
		p.vx = horizontal * p.speed
		p.x = p.x + p.vx
	}

	if vertical != 0 || horizontal != 0 {
		if isFire == false {
			p.deg = radToDeg(math.Atan2(vertical, horizontal))
		}
	}
}

func (p *Player) fire() {
	// TODO:
}
