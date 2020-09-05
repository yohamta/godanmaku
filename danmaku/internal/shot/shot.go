package shot

import (
	"math"

	"github.com/yohamta/godanmaku/danmaku/internal/effect"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
)

// Shot represents shooter
type Shot struct {
	x, y          float64
	width, height float64
	currField     *field.Field
	isActive      bool
	speed         float64
	vx            float64
	vy            float64
	degree        int
	spr           *sprite.Sprite
	sprIndex      int
}

// Kind represetns the kind of the shot
type Kind int

const (
	KindPlayerNormal Kind = iota
	KindEnemyNormal
)

// NewShot returns initialized struct
func NewShot(f *field.Field) *Shot {
	sh := &Shot{}
	sh.currField = f

	return sh
}

// Init inits the shot accoring to the kind
func (sh *Shot) Init(kind Kind, degree int) {
	sh.isActive = true

	switch kind {
	case KindPlayerNormal:
		sh.spr = sprite.PlayerBullet
		sh.setSize(4, 4)
		sh.setSpeed(2.56, degree)
		break
	case KindEnemyNormal:
		sh.spr = sprite.RandomEnemyShot()
		sh.setSize(10, 10)
		sh.setSpeed(1.44, degree)
	}
}

// IsActive returns if this is active
func (sh *Shot) IsActive() bool {
	return sh.isActive
}

// GetX returns x
func (sh *Shot) GetX() float64 {
	return sh.x
}

// GetY returns y
func (sh *Shot) GetY() float64 {
	return sh.y
}

// GetWidth returns width
func (sh *Shot) GetWidth() float64 {
	return sh.width
}

// GetHeight returns height
func (sh *Shot) GetHeight() float64 {
	return sh.height
}

// GetDegree returns the degree
func (sh *Shot) GetDegree() int {
	return sh.degree
}

// SetPosition sets the position
func (sh *Shot) SetPosition(x, y float64) {
	sh.x = x
	sh.y = y
}

// Draw draws this
func (sh *Shot) Draw(screen *ebiten.Image) {
	spr := sh.spr
	spr.SetPosition(sh.x, sh.y)
	spr.SetIndex(sh.sprIndex)
	spr.Draw(screen)
}

// Move moves this
func (sh *Shot) Move() {
	sh.SetPosition(sh.x+sh.vx, sh.y+sh.vy)
	if util.IsOutOfArea(sh, sh.currField) {
		sh.isActive = false
	}
}

// SetField returns field
func (sh *Shot) SetField(f *field.Field) {
	sh.currField = f
}

// OnHit should be called on hit something
func (sh *Shot) OnHit() {
	sh.isActive = false
	effect.CreateHitEffect(sh.x, sh.y)
}

func (sh *Shot) setSize(width, height float64) {
	sh.width = width
	sh.height = height
}

func (sh *Shot) setSpeed(speed float64, degree int) {
	sh.speed = speed
	sh.degree = degree
	sh.vx = math.Cos(util.DegToRad(sh.degree)) * speed
	sh.vy = math.Sin(util.DegToRad(sh.degree)) * speed
}
