package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/yohamta/godanmaku/danmaku"
)

type size struct {
	width  int
	height int
}

var (
	window size
	game   *danmaku.Game
)

func init() {
	game, err := danmaku.NewGame()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(game)
}

// SetWindowSize sets the window size
func SetWindowSize(width, height int) {
	window.width = width
	window.height = height
	game.SetWindowSize(width, height)
}

// dummy code for binding test

type Counter struct {
	Value int
}

func (c *Counter) Inc() { c.Value++ }

func NewCounter() *Counter { return &Counter{5} }

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
