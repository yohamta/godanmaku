package shooter

import (
	"math"
	"math/rand"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/effect"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
)

type PlayerController struct{}

func (c *PlayerController) init(sh *Shooter) {}

func (c *PlayerController) createLocusEffect(sh *Shooter, slottle float64) {
	if slottle < 0.5 {
		return
	}
	if sh.updateCount%int(10/slottle) == 0 {
		x, y := sh.GetPosition()
		effect.CreateLocusEffect(x, y)
	}
}

func (c *PlayerController) update(sh *Shooter) {
	x := sh.GetX()
	y := sh.GetY()
	f := sh.field

	vertical := shared.GameInput.Vertical
	horizontal := shared.GameInput.Horizontal

	slottle := math.Abs(vertical) + math.Abs(horizontal)
	c.createLocusEffect(sh, slottle)

	if vertical != 0 {
		sh.vy = vertical * sh.speed
		y = y + sh.vy
		y = math.Max(f.GetTop()+sh.GetHeight()/2, y)
		y = math.Min(f.GetBottom()-sh.GetHeight()/2, y)
	}

	if horizontal != 0 {
		sh.vx = horizontal * sh.speed
		x = x + sh.vx
		x = math.Max(f.GetLeft()+sh.GetWidth()/2, x)
		x = math.Min(f.GetRight()-sh.GetWidth()/2, x)
	}

	sh.SetPosition(x, y)

	if vertical != 0 || horizontal != 0 {
		if shared.GameInput.Fire == false {
			sh.degree = util.RadToDeg(math.Atan2(vertical, horizontal))
		}
	}
}

func (c *PlayerController) draw(sh *Shooter, screen *ebiten.Image) {
	sh.spr.SetPosition(sh.GetX()-shared.OffsetX, sh.GetY()-shared.OffsetY)
	sh.spr.SetIndex(util.DegreeToDirectionIndex(sh.degree))
	sh.spr.Draw(screen)
	healthBar.Draw(sh.x-shared.OffsetX, sh.y+sh.height/2-shared.OffsetY+5,
		float64(sh.life)/float64(sh.maxLife), screen)
}

func (c *PlayerController) isArrived(sh *Shooter) bool {
	return math.Abs(sh.y-sh.destination.y) < sh.GetHeight() &&
		math.Abs(sh.x-sh.destination.x) < sh.GetWidth()
}

func (c *PlayerController) updateDestination(sh *Shooter) {
	f := sh.field
	x := (f.GetRight() - f.GetLeft()) * rand.Float64()
	y := (f.GetBottom() - f.GetTop()) * rand.Float64()
	sh.destination.x = x
	sh.destination.y = y
	rad := math.Atan2(y-sh.y, x-sh.x)
	sh.vx = math.Cos(rad) * sh.speed
	sh.vy = math.Sin(rad) * sh.speed
}
