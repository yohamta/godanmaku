package effect

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	locusTTL  = 40.
	locusSize = 7.
)

type locus struct{ *baseController }

func (c *locus) init(e *Effect) {
}

func (c *locus) draw(e *Effect, screen *ebiten.Image) {
	if e.updateCount < e.waitFrame {
		return
	}
	if e.updateCount > locusTTL {
		return
	}

	size := (1. - float64(e.updateCount)/locusTTL) * locusSize * 1.3
	strength := 1. - float64(e.updateCount)/locusTTL

	c.drawGrowEffect(e, size, size, strength, screen)
}

func (c *locus) update(e *Effect) {
	if e.updateCount < e.waitFrame {
		return
	}
	if e.updateCount >= locusTTL {
		e.isActive = false
	}
}
