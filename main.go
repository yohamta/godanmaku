package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/miyahoyo/godanmaku/danmaku"
)

type size struct {
	width  int
	height int
}

func main() {
	window := &size{480, 800}

	ebiten.SetWindowTitle("Danmaku")
	ebiten.SetWindowSize(window.width, window.height)

	game, err := danmaku.NewGame()
	if err != nil {
		panic(err)
	}

	game.SetWindowSize(window.width, window.height)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
