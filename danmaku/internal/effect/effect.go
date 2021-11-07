package effect

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"
)

type Effect struct {
	x             float64
	y             float64
	isActive      bool
	controller    controller
	updateCount   int
	spriteFrame   int
	waitFrame     int
	callback      func()
	callbackFrame int
	scale         float64
	rotate        float64
	sprite        *ganim8.Sprite
	fps           int
	se            sound.SeKind
	sePlayed      bool
}

func NewEffect() *Effect {
	e := &Effect{}

	return e
}

func (e *Effect) IsActive() bool {
	return e.isActive
}

func (e *Effect) Draw(screen *ebiten.Image) {
	e.controller.draw(e, screen)
}

func (e *Effect) Update() {
	e.controller.update(e)
	e.updateCount++
}

func (e *Effect) init(c controller, x, y float64) {
	e.x = x
	e.y = y
	e.isActive = true
	e.controller = c
	e.updateCount = 0
	e.spriteFrame = 0
	e.waitFrame = 0
	e.scale = 1
	e.rotate = 0
	e.callback = nil
	e.callbackFrame = 0
	e.se = -1
	e.sePlayed = false
	c.init(e)
}
