package player

import (
	"bytes"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/actor"
	"github.com/yohamta/godanmaku/danmaku/internal/input"
	"github.com/yohamta/godanmaku/danmaku/internal/resources/images"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

const (
	initPlayerSpeed = 5.8
	downSpeed       = 1.2
)

// Player represents player of the game
type Player struct {
	sprite *sprite.Sprite
	actor  *actor.Actor
	vx     float64
	vy     float64
	degree int
}

// New returns initialized Player
func New() *Player {
	p := &Player{}

	img, _, _ := image.Decode(bytes.NewReader(images.PLAYER))
	sp := sprite.New(&img, 8)
	p.sprite = sp

	p.actor = &actor.Actor{}
	p.actor.SetPosition(120, 160)

	p.actor.SetSpeed(initPlayerSpeed)

	return p
}

// Draw draws this sprite
func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.SetPosition(p.actor.X, p.actor.Y)
	p.sprite.Draw(screen)
}

// Update updates the player's state
func (p *Player) Update(input *input.GameInput) {
	isMoving := false

	// Up
	if input.Up != 0 && (input.Right+input.Left == 0 || input.Right+input.Left == 2) {
		p.vx = 0
		p.vy = -p.actor.Speed
		p.actor.Y = p.actor.Y - p.actor.Speed

		if input.Fire == false {
			p.actor.SetDirection(8)
		}
		isMoving = true
	}

	// Down
	if input.Down != 0 && (input.Right+input.Left == 0 || input.Right+input.Left == 2) {
		p.vx = 0
		p.vy = p.actor.Speed
		p.actor.Y = p.actor.Y + p.actor.Speed

		if input.Fire == false {
			p.actor.SetDirection(2)
		}
		isMoving = true
	}

	// Left
	if input.Left != 0 && (input.Up+input.Down == 0 || input.Up+input.Down == 2) {
		p.vx = -p.actor.Speed
		p.vy = 0
		p.actor.X = p.actor.X - p.actor.Speed

		if input.Fire == false {
			p.actor.SetDirection(4)
		}
		isMoving = true
	}

	// Right
	if input.Right != 0 && (input.Up+input.Down == 0 || input.Up+input.Down == 2) {
		p.vx = p.actor.Speed
		p.vy = 0
		p.actor.X = p.actor.X + p.actor.Speed

		if input.Fire == false {
			p.actor.SetDirection(6)
		}
		isMoving = true
	}

	// Diagonal
	if isMoving == false {
		if input.Up != 0 && input.Right != 0 {
			p.vx = p.actor.NSpd
			p.vy = -p.actor.NSpd
			p.actor.X = p.actor.X + p.actor.NSpd
			p.actor.Y = p.actor.Y - p.actor.NSpd

			if input.Fire == false {
				p.actor.SetDirection(9)
			}
			isMoving = true
		}
		if input.Up != 0 && input.Left != 0 {
			p.vx = -p.actor.NSpd
			p.vy = -p.actor.NSpd
			p.actor.X = p.actor.X - p.actor.NSpd
			p.actor.Y = p.actor.Y - p.actor.NSpd

			if input.Fire == false {
				p.actor.SetDirection(7)
			}
			isMoving = true
		}
		if input.Down != 0 && input.Right != 0 {
			p.vx = p.actor.NSpd
			p.vy = p.actor.NSpd
			p.actor.X = p.actor.X + p.actor.NSpd
			p.actor.Y = p.actor.Y + p.actor.NSpd
			if input.Fire == false {
				p.actor.SetDirection(4)
			}
			isMoving = true
		}
		if input.Down != 0 && input.Left != 0 {
			p.vx = -p.actor.NSpd
			p.vy = p.actor.NSpd
			p.actor.X = p.actor.X - p.actor.NSpd
			p.actor.Y = p.actor.Y + p.actor.NSpd
			if input.Fire == false {
				p.actor.SetDirection(1)
			}
			isMoving = true
		}
	}

	// Inertia
	if isMoving == false {
		vx := p.vx
		vy := p.vy

		if vx > 0 {
			p.vx = math.Max(vx-downSpeed, 0)
		} else {
			p.vx = math.Min(vx+downSpeed, 0)
		}
		if vy > 0 {
			p.vy = math.Max(vy-downSpeed, 0)
		} else {
			p.vy = math.Min(vy+downSpeed, 0)
		}

		p.actor.X = p.actor.X + p.vx
		p.actor.Y = p.actor.Y + p.vy
	}

}
