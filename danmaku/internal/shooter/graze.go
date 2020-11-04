package shooter

import (
	"github.com/hajimehoshi/ebiten"
)

type GrazeController struct{}

func (c *GrazeController) init(sh *Shooter) {
}

func (c *GrazeController) update(sh *Shooter) {
	sh.SetPosition(sh.owner.GetPosition())
}

func (c *GrazeController) draw(sh *Shooter, screen *ebiten.Image) {
}
