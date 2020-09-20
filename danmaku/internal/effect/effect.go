package effect

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
)

// Effect represents the base of player, enemy, shots
type Effect struct {
	x           float64
	y           float64
	isActive    bool
	controller  controller
	updateCount int
	spriteFrame int
	waitFrame   int
	callback    func()
	scale       float64
	rotate      float64
}

// NewEffect creates new effect
func NewEffect() *Effect {
	e := &Effect{}

	return e
}

// IsActive returns if this is active
func (e *Effect) IsActive() bool {
	return e.isActive
}

// Draw draws the player
func (e *Effect) Draw(screen *ebiten.Image) {
	e.controller.draw(e, screen)
}

// Update updates the effect
func (e *Effect) Update() {
	e.controller.update(e)
	e.updateCount++
}

func (e *Effect) init(c controller, x, y float64) {
	e.x = x
	e.y = y
	e.isActive = true
	e.controller = c
	e.updateCount = 0
	e.spriteFrame = 0
	e.waitFrame = 0
	e.scale = 1
	e.rotate = 0
	c.init(e)
}

var (
	hitController       = &hit{newBaseControlelr()}
	explosionController = &explosion{newBaseControlelr()}
	jumpController      = &jump{newBaseControlelr()}
	locusController     = &locus{newBaseControlelr()}
)

// CreateLocusEffect creates an effect
func CreateLocusEffect(x, y float64) {
	e := (*Effect)(shared.BackEffects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(locusController, x, y)
	e.waitFrame = 3
}

// CreateHitEffect creates an effect
func CreateHitEffect(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(hitController, x, y)
}

// CreateExplosion creates an effect
func CreateExplosion(x, y float64) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(explosionController, x, y)
}

// CreateJump creates an effect
func CreateJump(x, y float64, wait int, callback func()) {
	e := (*Effect)(shared.Effects.CreateFromPool())
	if e == nil {
		return
	}
	e.init(jumpController, x, y)
	e.waitFrame = wait
	e.callback = callback
}
