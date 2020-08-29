package firebutton

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"

	"github.com/yohamta/godanmaku/danmaku/internal/ui/button"
)

// FireButton represents Button
type FireButton struct {
	button.Button
	color color.RGBA
}

// New returns new FireButton
func New() *FireButton {
	fButton := &FireButton{
		*(button.New(100, 50, 50, 40)),
		color.RGBA{0, 0xff, 0, 0xff},
	}
	return fButton
}

// Update updates the color of button
func (fButton *FireButton) Update() {
	if fButton.IsPressing() {
		fButton.color = color.RGBA{0xff, 0, 0, 0xff}
	} else {
		fButton.color = color.RGBA{0, 0xff, 0, 0xff}
	}
}

// Draw draws button to the screen
func (fButton *FireButton) Draw(screen *ebiten.Image) {
	x, y := fButton.GetPosition()
	w, h := fButton.GetSize()
	paint.FillRect(screen, paint.Rect{X: x, Y: y, W: w, H: h}, fButton.color)
}
