package player

import (
	"bytes"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/actor"
	"github.com/yohamta/godanmaku/danmaku/internal/resources/images"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

const (
	initPlayerSpeed = 2
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
	adjust := 22.5
	spriteIndex := int(float64(p.actor.Deg)+90.0+360.0+adjust) % 360 / 45
	p.sprite.SetIndex(spriteIndex)
	p.sprite.Draw(screen)
}

// Move moves player
func (p *Player) Move(horizontal float64, vertical float64, isFire bool) {
	if vertical != 0 {
		p.vy = vertical * p.actor.Speed
		p.actor.Y = p.actor.Y + p.vy
	}

	if horizontal != 0 {
		p.vx = horizontal * p.actor.Speed
		p.actor.X = p.actor.X + p.vx
	}

	if vertical != 0 || horizontal != 0 {
		degree := int(math.Atan2(vertical, horizontal) * 180 / math.Pi)
		if isFire == false {
			p.actor.SetDeg(degree)
		}
	}
}
