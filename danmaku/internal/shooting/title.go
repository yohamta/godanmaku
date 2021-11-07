package shooting

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
)

var flex *furex.Flex

func DrawTitle(screen *ebiten.Image) {
	if flex == nil {
		return
	}
	screen.Clear()
	flex.Draw(screen)
}

func UpdateTitle() {
	if flex == nil && isInitialized {
		flex = furex.NewFlex(screenSize.X, screenSize.Y)
		flex.Direction = furex.Column
		flex.Justify = furex.JustifyCenter
		flex.AlignItems = furex.AlignItemCenter
		flex.AlignContent = furex.AlignContentCenter

		startButton := NewStartButton()
		flex.AddChild(startButton)
	}
	flex.Update()
}
