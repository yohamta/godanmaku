package shooter

import (
	"math"

	"github.com/yohamta/godanmaku/danmaku/internal/effect"

	"github.com/yohamta/godanmaku/danmaku/internal/flyweight"

	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/weapon"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
)

// Target represents target
type Target interface {
	GetX() float64
	GetY() float64
	IsDead() bool
}

// Shooter represents shooter
type Shooter struct {
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
	life          int
	mainWeapon    weapon.Weapon
	target        Target
	movdweTo      struct{ x, y float64 }
	shotsPool     *flyweight.Pool
}

// NewShooter creates shooter struct
func NewShooter() *Shooter {
	sh := &Shooter{}

	return sh
}

// GetX returns x
func (sh *Shooter) GetX() float64 {
	return sh.x
}

// GetY returns y
func (sh *Shooter) GetY() float64 {
	return sh.y
}

// GetPosition returns the position
func (sh *Shooter) GetPosition() (float64, float64) {
	return sh.x, sh.y
}

// GetWidth returns width
func (sh *Shooter) GetWidth() float64 {
	return sh.width
}

// GetHeight returns height
func (sh *Shooter) GetHeight() float64 {
	return sh.height
}

// GetDegree returns height
func (sh *Shooter) GetDegree() int {
	return sh.degree
}

// IsActive returns if this is active
func (sh *Shooter) IsActive() bool {
	return sh.isActive
}

// GetMainSpriteIndex returns sprite
func (sh *Shooter) GetMainSpriteIndex() int {
	return sh.sprIndex
}

// SetMainSpriteIndex sets the sprite index
func (sh *Shooter) SetMainSpriteIndex(index int) {
	sh.sprIndex = index
}

// SetSpeed sets the speed
func (sh *Shooter) SetSpeed(speed float64, degree int) {
	sh.speed = speed
	sh.degree = degree
	sh.vx = math.Cos(util.DegToRad(sh.degree)) * speed
	sh.vy = math.Sin(util.DegToRad(sh.degree)) * speed
}

// AddDamage adds damage to this playe
func (sh *Shooter) AddDamage(damage int) {
	sh.life -= damage
	if sh.life <= 0 {
		sh.isActive = false
		effect.CreateExplosion(sh.x, sh.y)
	}
}

// IsDead returns if this is active
func (sh *Shooter) IsDead() bool {
	return sh.life <= 0
}

// SetWeapon adds damage to this playe
func (sh *Shooter) SetWeapon(mainWeapon weapon.Weapon) {
	sh.mainWeapon = mainWeapon
}

// SetPosition sets the position
func (sh *Shooter) SetPosition(x, y float64) {
	sh.x = x
	sh.y = y
}

// Fire fire the weapon
func (sh *Shooter) Fire() {
	sh.mainWeapon.Fire(sh.x, sh.y, sh.degree)
}

// SetTarget sets the target
func (sh *Shooter) SetTarget(target Target) {
	sh.target = target
}

func (sh *Shooter) setSize(width, height float64) {
	sh.width = width
	sh.height = height
}
