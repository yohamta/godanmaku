package shooter

import (
	"math"
	"math/rand"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
)

type FunnelController struct{}

func (c *FunnelController) init(sh *Shooter) {
	sh.funnelDegree = rand.Intn(359)
}

func (c *FunnelController) update(sh *Shooter) {
	if sh.updateCount%3 == 0 {
		sh.sprIndex++
		if sh.sprIndex >= sh.spr.Length() {
			sh.sprIndex = 0
		}
	}
	sh.funnelDegree += 2
	if sh.funnelDegree >= 360 {
		sh.funnelDegree = 0
	}
	sh.degree = sh.owner.degree
	sh.x = sh.owner.GetX() + math.Cos(util.DegToRad(sh.funnelDegree))*sh.GetWidth()*1.5
	sh.y = sh.owner.GetY() + math.Sin(util.DegToRad(sh.funnelDegree))*sh.GetWidth()*1.5
}

func (c *FunnelController) draw(sh *Shooter, screen *ebiten.Image) {
	sh.spr.SetPosition(sh.x-shared.OffsetX, sh.y-shared.OffsetY)
	sh.spr.SetIndex(sh.sprIndex)
	sh.spr.Draw(screen)
}
