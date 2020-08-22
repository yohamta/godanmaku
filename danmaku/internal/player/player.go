package player

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/actor"
	"github.com/yohamta/godanmaku/danmaku/internal/resources/images"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

// Player represents player of the game
type Player struct {
	sprite *sprite.Sprite
	actor  *actor.Actor
}

// New returns initialized Player
func New() *Player {
	p := &Player{}

	img, _, _ := image.Decode(bytes.NewReader(images.PLAYER))
	sp := sprite.New(&img, 8)
	p.sprite = sp

	p.actor = &actor.Actor{}
	p.actor.SetPosition(120, 160)

	return p
}

// Draw draws this sprite
func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.SetPosition(p.actor.X, p.actor.Y)
	p.sprite.Draw(screen)
}
