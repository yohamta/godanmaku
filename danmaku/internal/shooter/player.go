package shooter

import (
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
	"github.com/yohamta/godanmaku/danmaku/internal/weapon"
)

// Player represents player of the game
type Player struct {
	Shooter
	wep       weapon.Weapon
	shotSpeed float64
	shotSize  float64
}

// NewPlayer returns initialized Player
func NewPlayer() *Player {
	p := &Player{Shooter: *NewShooter()}

	p.life = 1
	p.SetSize(10, 10)
	p.SetPosition(120, 160)
	p.SetSpeed(2, 270)
	p.SetActive(true)

	return p
}

// Draw draws the player
func (p *Player) Draw(screen *ebiten.Image) {
	sprite.Player.SetPosition(p.GetX(), p.GetY())
	sprite.Player.SetIndex(util.DegreeToDirectionIndex(p.degree))
	sprite.Player.Draw(screen)
}

// Move moves the player
func (p *Player) Move(horizontal float64, vertical float64, isFire bool) {
	x := p.GetX()
	y := p.GetY()
	field := p.entity.GetField()

	if vertical != 0 {
		p.vy = vertical * p.speed
		y = y + p.vy
		y = math.Max(field.GetTop()+p.GetHeight()/2, y)
		y = math.Min(field.GetBottom()-p.GetHeight()/2, y)
	}

	if horizontal != 0 {
		p.vx = horizontal * p.speed
		x = x + p.vx
		x = math.Max(field.GetLeft()+p.GetWidth()/2, x)
		x = math.Min(field.GetRight()-p.GetWidth()/2, x)
	}

	p.SetPosition(x, y)

	if vertical != 0 || horizontal != 0 {
		if isFire == false {
			p.degree = util.RadToDeg(math.Atan2(vertical, horizontal))
		}
	}
}
