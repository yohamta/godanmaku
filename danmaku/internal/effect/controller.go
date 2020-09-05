package effect

import "github.com/hajimehoshi/ebiten"

// Controller represents effect controller
type Controller interface {
	init(e *Effect)
	update(e *Effect)
	draw(e *Effect, screen *ebiten.Image)
}

var (
	// Hit effect
	Hit = new(HitController)
	// Explosion effect
	Explosion = new(ExplosionController)
)
