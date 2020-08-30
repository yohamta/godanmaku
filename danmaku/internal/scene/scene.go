package scene

import (
	"github.com/hajimehoshi/ebiten"
)

// Scene represents scene interface
type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}
