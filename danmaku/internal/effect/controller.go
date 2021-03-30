package effect

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type controller interface {
	init(e *Effect)
	update(e *Effect)
	draw(e *Effect, screen *ebiten.Image)
}

type baseController struct{}

func newBaseControlelr() *baseController {
	c := &baseController{}
	return c
}

func (c *baseController) init(e *Effect) {}

func (c *baseController) draw(e *Effect, screen *ebiten.Image) {}

func (c *baseController) update(e *Effect) {}
