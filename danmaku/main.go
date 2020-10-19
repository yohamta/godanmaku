package danmaku

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shooting"
)

type Game struct {
	isInitialized bool
	screenSize    image.Point
}

type Scene interface {
	Update()
	Layout(width, height int)
	Draw(screen *ebiten.Image)
}

var (
	stg *shooting.Shooting
)

func NewGame() (*Game, error) {
	game := &Game{}

	return game, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.isInitialized == false {
		stg = shooting.NewShooting()
		stg.Layout(g.screenSize.X, g.screenSize.Y)
		g.isInitialized = true
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	height := float64(240) / float64(outsideWidth) * float64(outsideHeight)
	g.screenSize = image.Pt(240, int(height))
	if stg != nil {
		stg.Layout(g.screenSize.X, g.screenSize.Y)
	}
	return g.screenSize.X, g.screenSize.Y
}
