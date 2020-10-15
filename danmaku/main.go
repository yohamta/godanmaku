package danmaku

import (
	"github.com/yotahamada/godanmaku/danmaku/internal/battle"
	"github.com/yotahamada/godanmaku/danmaku/internal/sound"
	"github.com/yotahamada/godanmaku/danmaku/internal/ui"

	"github.com/hajimehoshi/ebiten"

	"github.com/yotahamada/godanmaku/danmaku/internal/paint"
	"github.com/yotahamada/godanmaku/danmaku/internal/sprite"
)

type Game struct{}

type Scene interface {
	Update()
	Layout(width, height int)
	Draw(screen *ebiten.Image)
}

const screenScale = 2

var (
	screenWidth   = 240
	screenHeight  int
	isInitialized = false
	btl           *battle.Battle
	state         State
)

func NewGame() (*Game, error) {
	game := &Game{}

	return game, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	if isInitialized == false {
		loadResources()

		state = StateBattle

		btl = battle.NewBattle()
		btl.Layout(screenWidth, screenHeight)

		isInitialized = true
	}

	switch state {
	case StateBattle:
		btl.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch state {
	case StateBattle:
		btl.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	screenWidth = outsideWidth / screenScale
	screenHeight = outsideHeight / screenScale

	ui.ScreenWidth = screenWidth
	ui.ScreenHeight = screenHeight

	if btl != nil {
		btl.Layout(screenWidth, screenHeight)
	}

	return screenWidth, screenHeight
}

func loadResources() {
	paint.LoadFonts()
	sprite.LoadSprites()
	sound.Load()
}
