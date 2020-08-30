package shooting

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/actors"
)

var (
	myPlayer     *actors.Player
	myInput      *Input
	screenWidth  = 0
	screenHeight = 0
)

// Shooting represents shooting scene
type Shooting struct{}

// NewShootingOptions represents options for New func
type NewShootingOptions struct {
	ScreenWidth  int
	ScreenHeight int
}

// NewShooting returns new Shooting struct
func NewShooting(options NewShootingOptions) *Shooting {
	shooting := &Shooting{}

	screenWidth = options.ScreenWidth
	screenHeight = options.ScreenHeight

	myPlayer = actors.NewPlayer()
	myInput = NewInput(screenWidth, screenHeight)

	return shooting
}

// Update updates the scene
func (shooting *Shooting) Update() {
	myInput.Update()
	myPlayer.Move(myInput.Horizontal, myInput.Vertical, myInput.Fire)

}

// Draw draws the scene
func (shooting *Shooting) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	myPlayer.Draw(screen)
	myInput.Draw(screen)
}
