package danmaku

import (
	"image"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/shooting"
)

type Game struct{}

type Scene interface {
	Update()
	Layout(width, height int)
	Draw(screen *ebiten.Image)
}

const (
	widthAsDots = 240.
)

var (
	stg             *shooting.Shooting
	isWindowSizeSet = false
	isInitialized   bool
	screenSize      image.Point
)

func NewGame() (*Game, error) {
	game := &Game{}

	return game, nil
}

func (g *Game) Update() error {
	if isWindowSizeSet && isInitialized == false {
		stg = shooting.NewShooting()
		stg.Layout(screenSize.X, screenSize.Y)
		isInitialized = true
		return nil
	}
	stg.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if stg != nil {
		stg.Draw(screen)
	}
}

func (g *Game) SetWindowSize(width, height int) {
	screenSize.X = int(widthAsDots)
	screenSize.Y = int(widthAsDots / float64(width) * float64(height))
	isWindowSizeSet = true
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	if stg != nil {
		stg.Layout(screenSize.X, screenSize.Y)
	}
	return screenSize.X, screenSize.Y
}
