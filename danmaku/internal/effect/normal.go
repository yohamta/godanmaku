package effect

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miyahoyo/godanmaku/danmaku/internal/shared"
	"github.com/miyahoyo/godanmaku/danmaku/internal/sound"
)

const fps = 60

type normal struct{ *baseController }

func (c *normal) init(e *Effect) {}

func (c *normal) update(e *Effect) {
	if e.isActive == false {
		return
	}
	if e.updateCount < e.waitFrame {
		return
	}
	if e.updateCount%(fps/e.fps) == 0 {
		e.spriteFrame++
	}
	if e.updateCount >= e.waitFrame && e.se != -1 && !e.sePlayed {
		sound.PlaySe(e.se)
		e.sePlayed = true
	}
	if e.callbackFrame == e.spriteFrame {
		if e.callback != nil {
			e.callback()
			e.callback = nil
		}
	}
	if e.spriteFrame >= e.sprite.Length() {
		e.isActive = false
		if e.callback != nil {
			e.callback()
			e.callback = nil
		}
	}
}

func (c *normal) draw(e *Effect, screen *ebiten.Image) {
	if e.isActive == false {
		return
	}
	if e.updateCount < e.waitFrame {
		return
	}
	if e.spriteFrame >= e.sprite.Length() {
		return
	}
	e.sprite.SetIndex(e.spriteFrame)
	e.sprite.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	e.sprite.DrawWithScale(screen, e.scale)
}
