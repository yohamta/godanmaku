package shooter

import (
	"math"
	"math/rand"

	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/flyweight"

	"github.com/yohamta/godanmaku/danmaku/internal/shot"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
	"github.com/yohamta/godanmaku/danmaku/internal/weapon"
)

// Enemy represents enemy of the game
type Enemy struct {
	Shooter
	moveTo    struct{ x, y float64 }
	wep       weapon.Weapon
	shotSpeed float64
	shotSize  float64
}

// NewEnemy returns initialized Enemy
func NewEnemy(f *field.Field, shotsPool *flyweight.Pool) *Enemy {
	e := &Enemy{Shooter: *NewShooter()}
	e.field = f
	e.shotsPool = shotsPool

	return e
}

// Init inits the enemy
func (e *Enemy) Init(x, y float64) {
	width := 24.
	height := 24.
	e.setSize(width, height)
	e.SetPosition(x, y)
	e.SetSpeed(0.96, 90)
	e.SetWeapon(weapon.Normal(shot.EnemyShot, false))

	e.life = 3
	e.maxLife = e.life
	e.isActive = true
	e.spr = sprite.Enemy1
	e.updateMoveTo()
}

// Draw draws the enemy
func (e *Enemy) Draw(screen *ebiten.Image) {
	sprite.Enemy1.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Enemy1.SetIndex(util.DegreeToDirectionIndex(e.degree))
	sprite.Enemy1.Draw(screen)
}

// Update moves the enemy
func (e *Enemy) Update() {
	e.SetPosition(e.x+e.vx, e.y+e.vy)

	if e.isArrived() {
		e.updateMoveTo()
	}

	target := e.target

	if rand.Float64() < 0.05 {
		e.degree = util.RadToDeg(math.Atan2(target.GetY()-e.y, target.GetX()-e.x))
	}
}

func (e *Enemy) isArrived() bool {
	return math.Abs(e.y-e.moveTo.y) < e.GetHeight() &&
		math.Abs(e.x-e.moveTo.x) < e.GetWidth()
}

func (e *Enemy) updateMoveTo() {
	f := e.field
	x := (f.GetRight() - f.GetLeft()) * rand.Float64()
	y := (f.GetBottom() - f.GetTop()) * rand.Float64()
	e.moveTo.x = x
	e.moveTo.y = y
	rad := math.Atan2(y-e.y, x-e.x)
	e.vx = math.Cos(rad) * e.speed
	e.vy = math.Sin(rad) * e.speed
}
