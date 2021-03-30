package shot

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type controller interface {
	init(s *Shot)
	update(s *Shot)
	draw(s *Shot, screen *ebiten.Image)
}

type baseController struct{}

func (c *baseController) init(s *Shot) {}

func (c *baseController) update(s *Shot) {}

func (c *baseController) draw(s *Shot, screen *ebiten.Image) {}
