package shot

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/entity"
	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
)

// Shot represents shooter
type Shot struct {
	entity          *entity.Entity
	speed           float64
	vx              float64
	vy              float64
	degree          int
	mainSprite      *sprite.Sprite
	mainSpriteIndex int
}

// Kind represetns the kind of the shot
type Kind int

const (
	KindPlayerNormal Kind = iota
	KindEnemyNormal
)

// NewShot returns initialized struct
func NewShot() *Shot {
	sh := &Shot{entity: entity.NewEntity()}

	return sh
}

// Init inits the shot accoring to the kind
func (sh *Shot) Init(kind Kind, degree int) {
	sh.SetActive(true)

	switch kind {
	case KindPlayerNormal:
		sh.SetMainSprite(sprite.PlayerBullet)
		sh.SetSize(4, 4)
		sh.SetSpeed(2.56, degree)
		break
	case KindEnemyNormal:
		sh.SetMainSprite(sprite.RandomEnemyShot())
		sh.SetSize(10, 10)
		sh.SetSpeed(1.44, degree)
	}
}

// IsActive returns if this entity is active
func (sh *Shot) IsActive() bool {
	return sh.entity.IsActive()
}

// GetX returns x
func (sh *Shot) GetX() float64 {
	return sh.entity.GetX()
}

// GetY returns y
func (sh *Shot) GetY() float64 {
	return sh.entity.GetY()
}

// SetActive sets if this entity is active
func (sh *Shot) SetActive(isActive bool) {
	sh.entity.SetActive(isActive)
}

// GetMainSprite returns sprite
func (sh *Shot) GetMainSprite() *sprite.Sprite {
	return sh.mainSprite
}

// SetMainSprite sets the sprite
func (sh *Shot) SetMainSprite(mainSprite *sprite.Sprite) {
	sh.mainSprite = mainSprite
}

// GetMainSpriteIndex returns sprite
func (sh *Shot) GetMainSpriteIndex() int {
	return sh.mainSpriteIndex
}

// SetMainSpriteIndex sets the sprite index
func (sh *Shot) SetMainSpriteIndex(index int) {
	sh.mainSpriteIndex = index
}

// SetSpeed sets the speed
func (sh *Shot) SetSpeed(speed float64, degree int) {
	sh.speed = speed
	sh.degree = degree
	sh.vx = math.Cos(util.DegToRad(sh.degree)) * speed
	sh.vy = math.Sin(util.DegToRad(sh.degree)) * speed
}

// GetEntity returns the degree
func (sh *Shot) GetEntity() *entity.Entity {
	return sh.entity
}

// GetDegree returns the degree
func (sh *Shot) GetDegree() int {
	return sh.degree
}

// SetSize returns the size
func (sh *Shot) SetSize(width, height float64) {
	sh.entity.SetSize(width, height)
}

// SetPosition sets the position
func (sh *Shot) SetPosition(x, y float64) {
	sh.entity.SetPosition(x, y)
}

// Draw draws this
func (sh *Shot) Draw(screen *ebiten.Image) {
	sh.GetMainSprite().SetPosition(sh.entity.GetX(), sh.entity.GetY())
	sh.GetMainSprite().SetIndex(sh.GetMainSpriteIndex())
	sh.GetMainSprite().Draw(screen)
}

// Move moves this
func (sh *Shot) Move() {
	sh.entity.SetPosition(sh.entity.GetX()+sh.vx, sh.entity.GetY()+sh.vy)
	if sh.entity.IsOutOfField() {
		sh.SetActive(false)
	}
}

// SetField returns field
func (sh *Shot) SetField(f *field.Field) {
	sh.entity.SetField(f)
}
