package shooter

import (
	"math"

	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/shot"
	"github.com/yohamta/godanmaku/danmaku/internal/weapon"

	"github.com/yohamta/godanmaku/danmaku/internal/entity"
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
	entity          *entity.Entity
	speed           float64
	vx              float64
	vy              float64
	degree          int
	mainSprite      *sprite.Sprite
	mainSpriteIndex int
	life            int
	mainWeapon      weapon.Weapon
	target          Target
	movdweTo        struct{ x, y float64 }
}

// NewShooter creates shooter struct
func NewShooter() *Shooter {
	sh := &Shooter{entity: entity.NewEntity()}

	return sh
}

// GetX returns x
func (sh *Shooter) GetX() float64 {
	return sh.entity.GetX()
}

// GetY returns y
func (sh *Shooter) GetY() float64 {
	return sh.entity.GetY()
}

// GetPosition returns the position
func (sh *Shooter) GetPosition() (float64, float64) {
	return sh.entity.GetPosition()
}

// GetWidth returns width
func (sh *Shooter) GetWidth() float64 {
	return sh.entity.GetWidth()
}

// GetHeight returns height
func (sh *Shooter) GetHeight() float64 {
	return sh.entity.GetHeight()
}

// GetDegree returns height
func (sh *Shooter) GetDegree() int {
	return sh.degree
}

// IsActive returns if this entity is active
func (sh *Shooter) IsActive() bool {
	return sh.entity.IsActive()
}

// SetActive sets if this entity is active
func (sh *Shooter) SetActive(isActive bool) {
	sh.entity.SetActive(isActive)
}

// GetMainSprite returns sprite
func (sh *Shooter) GetMainSprite() *sprite.Sprite {
	return sh.mainSprite
}

// SetMainSprite sets the sprite
func (sh *Shooter) SetMainSprite(mainSprite *sprite.Sprite) {
	sh.mainSprite = mainSprite
}

// GetMainSpriteIndex returns sprite
func (sh *Shooter) GetMainSpriteIndex() int {
	return sh.mainSpriteIndex
}

// SetMainSpriteIndex sets the sprite index
func (sh *Shooter) SetMainSpriteIndex(index int) {
	sh.mainSpriteIndex = index
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
}

// IsDead returns if this is active
func (sh *Shooter) IsDead() bool {
	return sh.life <= 0
}

// SetMainWeapon adds damage to this playe
func (sh *Shooter) SetMainWeapon(mainWeapon weapon.Weapon) {
	sh.mainWeapon = mainWeapon
}

// SetSize returns the size
func (sh *Shooter) SetSize(width, height float64) {
	sh.entity.SetSize(width, height)
}

// SetPosition sets the position
func (sh *Shooter) SetPosition(x, y float64) {
	sh.entity.SetPosition(x, y)
}

// FireWeapon fire the weapon
func (sh *Shooter) FireWeapon(shots []*shot.Shot) {
	sh.mainWeapon.Fire(sh, shots)
}

// IsCollideWith check collistion with other
func (sh *Shooter) IsCollideWith(other *entity.Entity) bool {
	return sh.entity.IsCollideWith(other)
}

// SetField returns field
func (sh *Shooter) SetField(f *field.Field) {
	sh.entity.SetField(f)
}

// SetTarget sets the target
func (sh *Shooter) SetTarget(target Target) {
	sh.target = target
}
