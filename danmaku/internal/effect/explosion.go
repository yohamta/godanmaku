package effect

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

// ExplosionController represents player of the game
type ExplosionController struct{}

func (c *ExplosionController) init(e *Effect) {}

func (c *ExplosionController) draw(e *Effect, screen *ebiten.Image) {
	if e.spriteFrame >= sprite.Explosion.Length() {
		return
	}
	sprite.Explosion.SetIndex(e.spriteFrame)
	sprite.Explosion.SetPosition(e.x, e.y)
	sprite.Explosion.Draw(screen)
}

func (c *ExplosionController) update(e *Effect) {
	if e.updateCount > 0 && e.updateCount%2 == 0 {
		e.spriteFrame++
	}
	if e.spriteFrame >= sprite.Explosion.Length() {
		e.isActive = false
	}
}
