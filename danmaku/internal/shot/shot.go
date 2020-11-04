package shot

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/effect"
	"github.com/yohamta/godanmaku/danmaku/internal/quadtree"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
)

type point struct {
	x, y float64
}

var (
	laserCollisionIDMap = map[int]string{}
	laserAdjustMap      = map[int]*point{}
)

func init() {
	for i := 0; i <= 345; i += 15 {
		laserCollisionIDMap[i] = fmt.Sprintf("laser1_%d", i)
		x := 0.
		y := 0.
		if i <= 75 {
			x = 8
			y = 6
		} else if i <= 175 {
			x = -6
			y = 8
		} else if i <= 255 {
			x = -8
			y = -6
		} else {
			x = 8
			y = -6
		}
		laserAdjustMap[i] = &point{x, y}
	}
}

type Shooter interface {
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
	IsDead() bool
}

type Weapon interface {
	Fire(shooter Shooter, x, y float64, degree int)
}

type Shot struct {
	controller    controller
	x, y          float64
	width, height float64
	field         *field.Field
	isActive      bool
	speed         float64
	vx            float64
	vy            float64
	degree        int
	spr           *sprite.Sprite
	sprIndex      int
	updateCount   int
	quadNode      *quadtree.Node
	collisionBox  []*collision.Box
	shooter       Shooter
	funnelWeapon  Weapon
	itemKind      ItemKind
	isGrazed      bool
}

func NewShot(f *field.Field) *Shot {
	s := &Shot{}
	s.field = f

	return s
}

func (s *Shot) GetQuadtreeNode() *quadtree.Node {
	return s.quadNode
}

func (s *Shot) SetQuadtreeNode(n *quadtree.Node) {
	s.quadNode = n
}

func (s *Shot) IsActive() bool {
	return s.isActive
}

func (s *Shot) GetX() float64 {
	return s.x
}

func (s *Shot) GetY() float64 {
	return s.y
}

func (s *Shot) GetWidth() float64 {
	return s.width
}

func (s *Shot) GetHeight() float64 {
	return s.height
}

// Implement quadtree.Collider interface
func (s *Shot) GetRect() (x0 float64, y0 float64, x1 float64, y1 float64) {
	return s.x - s.width/2, s.y - s.height/2, s.x + s.width/2, s.y + s.height/2
}

func (s *Shot) GetCollisionBox() []*collision.Box {
	return s.collisionBox
}

func (s *Shot) Draw(screen *ebiten.Image) {
	s.controller.draw(s, screen)
}

func (s *Shot) Update() {
	s.updateCount++
	s.setPosition(s.x+s.vx, s.y+s.vy)
	if util.IsOutOfAreaEnoughly(s, s.field) {
		s.isActive = false
	}
	s.controller.update(s)
}

func (s *Shot) OnHit() {
	s.isActive = false
	if rand.Float64() > 0.5 {
		effect.CreateHitEffect(s.x, s.y)
	} else {
		effect.CreateHitLargeEffect(s.x, s.y)
	}
}

func (s *Shot) GetItemKind() ItemKind {
	return s.itemKind
}

func (s *Shot) SetInactive() {
	s.isActive = false
}

func (s *Shot) SetGrazed() {
	s.isGrazed = true
}

func (s *Shot) IsGrazed() bool {
	return s.isGrazed
}

func (s *Shot) setSize(width, height float64) {
	s.width = width
	s.height = height
}

func (s *Shot) setSpeed(speed float64, degree int) {
	s.speed = speed
	s.degree = degree
	s.vx = math.Cos(util.DegToRad(s.degree)) * speed
	s.vy = math.Sin(util.DegToRad(s.degree)) * speed
}

func (s *Shot) init(controller controller, shooter Shooter, x, y float64, degree int) {
	s.isActive = true
	s.x = x
	s.y = y
	s.degree = degree
	s.updateCount = 0
	s.controller = controller
	s.shooter = shooter
	s.isGrazed = false
	controller.init(s)
}

func (s *Shot) setPosition(x, y float64) {
	s.x = x
	s.y = y
}
