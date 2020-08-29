package danmaku

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/scene"
	"github.com/yohamta/godanmaku/danmaku/internal/scene/shooting"
)

var (
	currentScene  scene.Scene = nil
	screenWidth               = 240
	screenHeight              = 320
	isLayouted                = false
	isInitialized             = false
)

// Game implements ebiten.Game interface.
type Game struct {
}

// Update updates a game by one tick. The given argument represents a screen image.
func (g *Game) Update(screen *ebiten.Image) error {
	if isLayouted && !isInitialized {
		currentScene = shooting.New(shooting.NewOptions(shooting.NewOptions{
			ScreenWidth:  screenWidth,
			ScreenHeight: screenHeight,
		}))
		isInitialized = true
		return nil
	}
	if currentScene != nil {
		currentScene.Update()
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	currentScene.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	screenHeight = int(float64(screenWidth) / float64(outsideWidth) * float64(outsideHeight))
	isLayouted = true
	return screenWidth, screenHeight
}

// NewGame creates a game struct
func NewGame() (*Game, error) {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Danmaku")

	game := &Game{}

	return game, nil
}
