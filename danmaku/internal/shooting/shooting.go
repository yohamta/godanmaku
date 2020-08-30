package shooting

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shooting/actors"
	"github.com/yohamta/godanmaku/danmaku/internal/shooting/fields"
	"github.com/yohamta/godanmaku/danmaku/internal/shooting/inputs"
	"github.com/yohamta/godanmaku/danmaku/internal/shooting/weapons"
)

const (
	maxPlayerBulletsNum = 80
)

// PlayerWeapon represents interface of Player Weapon
type PlayerWeapon interface {
	Shot(x, y float64, degree int, playerShots []*actors.PlayerBullet)
}

var (
	input        *inputs.Input
	screenWidth  = 0
	screenHeight = 0

	field        *fields.Field
	player       *actors.Player
	playerShots  [maxPlayerBulletsNum]*actors.PlayerBullet
	playerWeapon PlayerWeapon
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

	input = inputs.NewInput(screenWidth, screenHeight)
	field = fields.NewField()

	// init actors
	player = actors.NewPlayer()
	for i := 0; i < maxPlayerBulletsNum; i++ {
		playerShots[i] = actors.NewPlayerShot()
	}

	playerWeapon = &weapons.PlayerWeapon1{}

	return stg
}

// Update updates the scene
func (stg *Shooting) Update() {
	input.Update()
	player.Move(input.Horizontal, input.Vertical, input.Fire, field)
	if input.Fire {
		x, y := player.GetPosition()
		playerWeapon.Shot(x, y, player.GetDeg(), playerShots[:])
	}

	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Move(field)
	}
}

// Draw draws the scene
func (stg *Shooting) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})

	field.Draw(screen)
	player.Draw(screen)
	input.Draw(screen)

	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Draw(screen)
	}

}
