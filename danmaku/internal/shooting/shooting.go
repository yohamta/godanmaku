package shooting

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	maxPlayerBulletsNum = 80
)

var (
	myInput      *Input
	screenWidth  = 0
	screenHeight = 0

	// actors
	myPlayer      *Player
	playerBullets [maxPlayerBulletsNum]*PlayerBullet
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
	stg := &Shooting{}

	screenWidth = options.ScreenWidth
	screenHeight = options.ScreenHeight

	myInput = NewInput(screenWidth, screenHeight)

	// init actors
	myPlayer = NewPlayer()
	for i := 0; i < maxPlayerBulletsNum; i++ {
		playerBullets[i] = NewPlayerBullet()
	}

	return stg
}

// Update updates the scene
func (stg *Shooting) Update() {
	myInput.update()
	myPlayer.move(myInput.Horizontal, myInput.Vertical, myInput.Fire)
	myPlayer.fire()
}

// Draw draws the scene
func (stg *Shooting) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	myPlayer.draw(screen)
	myInput.draw(screen)
}
