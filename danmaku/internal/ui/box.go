package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
)

// Box represents colored box to render on screen
type Box struct {
	position    position
	size        size
	color       color.Color
	offsetImage *ebiten.Image
}

// NewBox returns new created Box
func NewBox(x, y, width, height int, clr color.Color) *Box {
	b := &Box{}
	b.position.x = x
	b.position.y = y
	b.size.w = width
	b.size.h = height
	b.color = clr

	offsetImage, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	paint.FillRect(offsetImage, paint.Rect{X: 0, Y: 0, W: width, H: height}, clr)
	b.offsetImage = offsetImage

	return b
}

// Draw draws box to the screen
func (b *Box) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.position.x), float64(b.position.y))
	screen.DrawImage(b.offsetImage, op)
}
