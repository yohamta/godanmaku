package shooting

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	maxPlayerBulletsNum = 80
)

var (
	input        *Input
	screenWidth  = 0
	screenHeight = 0

	field       *Field
	player      *Player
	playerShots [maxPlayerBulletsNum]*PlayerBullet
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

	input = NewInput(screenWidth, screenHeight)
	field = NewField()

	// init actors
	player = NewPlayer()
	for i := 0; i < maxPlayerBulletsNum; i++ {
		playerShots[i] = NewPlayerShot()
	}

	return stg
}

// Update updates the scene
func (stg *Shooting) Update() {
	input.update()
	player.move(input.Horizontal, input.Vertical, input.Fire, field)
	if input.Fire {
		player.weapon.shot()
	}

	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.isActive == false {
			continue
		}
		p.move(field)
	}
}

// Draw draws the scene
func (stg *Shooting) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})

	field.draw(screen)
	player.draw(screen)
	input.draw(screen)

	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.isActive == false {
			continue
		}
		p.draw(screen)
	}

}
