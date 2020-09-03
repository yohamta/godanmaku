package shooter

import (
	"math"
	"math/rand"

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
func NewEnemy() *Enemy {
	e := &Enemy{Shooter: *NewShooter()}
	return e
}

// Init inits the enemy
func (e *Enemy) Init() {
	field := e.entity.GetField()
	fieldWidth := field.GetRight() - field.GetLeft()

	width := 8.
	height := 8.
	e.SetSize(width, height)
	e.SetPosition(rand.Float64()*float64(fieldWidth-width)+float64(width/2), 30)
	e.SetSpeed(0.96, 90)
	e.SetMainWeapon(weapon.NewNormal(shot.KindEnemyNormal))

	e.life = 1
	e.SetActive(true)
	e.SetMainSprite(sprite.Enemy1)
	e.updateMoveTo()
}

// Draw draws the enemy
func (e *Enemy) Draw(screen *ebiten.Image) {
	sprite.Enemy1.SetPosition(e.GetX(), e.GetY())
	sprite.Enemy1.SetIndex(util.DegreeToDirectionIndex(e.degree))
	sprite.Enemy1.Draw(screen)
}

// Move moves the enemy
func (e *Enemy) Move() {
	e.SetPosition(e.GetX()+e.vx, e.GetY()+e.vy)

	if e.isArrived() {
		e.updateMoveTo()
	}

	target := e.target
	e.degree = util.RadToDeg(math.Atan2(target.GetY()-e.GetY(), target.GetX()-e.GetX()))
}

func (e *Enemy) isArrived() bool {
	return math.Abs(e.GetY()-e.moveTo.y) < e.GetHeight() &&
		math.Abs(e.GetX()-e.moveTo.x) < e.GetWidth()
}

func (e *Enemy) updateMoveTo() {
	field := e.entity.GetField()
	x := (field.GetRight() - field.GetLeft()) * rand.Float64()
	y := (field.GetBottom() - field.GetTop()) * rand.Float64()
	e.moveTo.x = x
	e.moveTo.y = y
	rad := math.Atan2(y-e.GetY(), x-e.GetX())
	e.vx = math.Cos(rad) * e.speed
	e.vy = math.Sin(rad) * e.speed
}
