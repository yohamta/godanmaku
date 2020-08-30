package actors

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

// EnemyKind represends the kind of enemy
type EnemyKind int

const (
	EnemyKindBall EnemyKind = iota
)

// Enemy represents enemy of the game
type Enemy struct {
	Actor
	moveTo   position
	isActive bool
	life     int
	point    int
}

// NewEnemy returns initialized Enemy
func NewEnemy() *Enemy {
	actor := &Actor{}
	e := &Enemy{Actor: *actor}
	e.isActive = false
	return e
}

// InitEnemy inits the enemy
func (e *Enemy) InitEnemy(kind EnemyKind) {
	fieldWidth := boundarizer.GetRight() - boundarizer.GetLeft()
	switch kind {
	case EnemyKindBall:
		e.width = 8
		e.height = 8
		e.speed = 3
		e.deg = 90
		e.speed = 0.96
		e.x = rand.Float64()*float64(fieldWidth-e.width) + float64(e.width/2)
		e.y = 30
		e.life = 1
		e.point = 100
		e.isActive = true
		e.updateMoveTo()
	}
}

// Draw draws the enemy
func (e *Enemy) Draw(screen *ebiten.Image) {
	sprite.Enemy1.SetPosition(e.x, e.y)
	sprite.Enemy1.SetIndex(degreeToDirectionIndex(e.deg))
	sprite.Enemy1.Draw(screen)
}

// Move moves the enemy
func (e *Enemy) Move(player *Player) {
	e.x = e.x + e.vx
	e.y = e.y + e.vy

	if e.isArrived() {
		e.updateMoveTo()
	}

	rad := math.Atan2(player.x-e.x, player.y-e.y)
	e.deg = radToDeg(rad)
}

func (e *Enemy) isArrived() bool {
	return int(math.Abs(e.y-e.moveTo.y)) < e.height+10 &&
		int(math.Abs(e.x-e.moveTo.x)) < e.width+10
}

// IsActive returns if this is active
func (e *Enemy) IsActive() bool {
	return e.isActive
}

// IsDead returns if this is active
func (e *Enemy) IsDead() bool {
	return e.life <= 0
}

func (e *Enemy) updateMoveTo() {
	x, y := getRandomLocation()
	e.moveTo.x = x
	e.moveTo.y = y
	rad := math.Atan2(y-e.y, x-e.x)
	e.vx = math.Cos(rad) * e.speed
	e.vy = math.Sin(rad) * e.speed
}

func getRandomLocation() (float64, float64) {
	x := float64(boundarizer.GetRight()-boundarizer.GetLeft()) * rand.Float64()
	y := float64(boundarizer.GetBottom()-boundarizer.GetTop()) * rand.Float64()
	return x, y
}

// AddDamage adds damage to this enemy
func (e *Enemy) AddDamage(damage int) {
	e.life -= damage
	if e.IsDead() {
		e.isActive = false
	}
}
