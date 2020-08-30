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
	actor.Actor
	sprite *sprite.Sprite
	vx     float64
	vy     float64
	degree int
}

// New returns initialized Player
func New() *Player {
	actor := &actor.Actor{}
	p := &Player{Actor: *actor}

	img, _, _ := image.Decode(bytes.NewReader(images.PLAYER))
	sp := sprite.New(&img, 8)
	p.sprite = sp

	p.SetPosition(120, 160)

	p.SetSpeed(initPlayerSpeed)

	return p
}

// Draw draws this sprite
func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.SetPosition(p.X, p.Y)
	adjust := 22.5
	spriteIndex := int(float64(p.Deg)+90.0+360.0+adjust) % 360 / 45
	p.sprite.SetIndex(spriteIndex)
	p.sprite.Draw(screen)
}

// Move moves player
func (p *Player) Move(horizontal float64, vertical float64, isFire bool) {
	if vertical != 0 {
		p.vy = vertical * p.Speed
		p.Y = p.Y + p.vy
	}

	if horizontal != 0 {
		p.vx = horizontal * p.Speed
		p.X = p.X + p.vx
	}

	if vertical != 0 || horizontal != 0 {
		degree := int(math.Atan2(vertical, horizontal) * 180 / math.Pi)
		if isFire == false {
			p.SetDeg(degree)
		}
	}
}
