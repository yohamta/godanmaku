package shooter

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type GrazeController struct{}

func (c *GrazeController) init(sh *Shooter) {
}

func (c *GrazeController) update(sh *Shooter) {
	sh.SetPosition(sh.owner.GetPosition())
}

func (c *GrazeController) draw(sh *Shooter, screen *ebiten.Image) {
}
