package danmaku

import (
	"image/color"

	"github.com/yohamta/godanmaku/danmaku/internal/ui"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/paint"
	"github.com/yohamta/godanmaku/danmaku/internal/shooting"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

// Game implements ebiten.Game interface.
type Game struct {
	lcnt int
}

// Scene represents scene interface
type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

var (
	screenWidth     = 240
	screenHeight    = 320
	isInitialized   = false
	isWindowSizeSet = false
)

// Update updates a game by one tick. The given argument represents a screen image.
func (g *Game) Update(screen *ebiten.Image) error {
	if isWindowSizeSet && !isInitialized {
		paint.LoadFonts()
		sprite.LoadSprites()
		ui.SetScreenSize(screenWidth, screenHeight)
		ui.SetRootView(shooting.NewShooting())
		isInitialized = true
		return nil
	}

	ui.Update()

	return nil
}

// SetWindowSize inits the gaem
func (g *Game) SetWindowSize(width, height int) {
	screenHeight = int(float64(screenWidth) / float64(width) * float64(height))
	isWindowSizeSet = true
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	ui.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// NewGame creates a game struct
func NewGame() (*Game, error) {
	game := &Game{}
	game.lcnt = 0

	return game, nil
}
