package shooting

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"
)

type StartButton struct {
	isPressing bool
}

func NewStartButton() *StartButton {
	return new(StartButton)
}

func (b *StartButton) Size() (int, int) {
	return 80, 30
}

func (b *StartButton) HandlePress(x, y int) {
	b.isPressing = true
	sound.Load()
}

func (b *StartButton) HandleRelease(x, y int, isCancel bool) {
	b.isPressing = false
	if isCancel == false {
		startGame()
	}
}

func (b *StartButton) Draw(screen *ebiten.Image, frame image.Rectangle) {
	if b.isPressing {
		paint.FillRect(screen, frame, color.RGBA{
			0xb4, 0x20, 0x2a, 0xff,
		})
	} else {
		paint.DrawRect(screen, frame, color.White, 1)
	}
	paint.DrawTextWithOptions(screen, "START", (frame.Min.X+frame.Max.X)/2.,
		(frame.Min.Y+frame.Max.Y)/2., paint.DrawTextOptions{
			Color:    color.White,
			Width:    frame.Max.X - frame.Min.X,
			HAligh:   paint.HAlignCenter,
			FontSize: paint.FontSizeMedium,
		})
}
