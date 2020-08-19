package main

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku"
)

func main() {
	game, err := danmaku.NewGame()
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
