package ui

import (
	"image/color"

	"github.com/yotahamada/godanmaku/danmaku/internal/uikit"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/godanmaku/danmaku/internal/paint"
)

type Rectangle struct {
	uikit.ViewEventHandlerFuncs

	w, h        int
	color       color.Color
	offsetImage *ebiten.Image
}

func NewRectangle(w, h int, clr color.Color) *Rectangle {
	r := new(Rectangle)

	r.w = w
	r.h = h
	r.color = clr

	return r
}

func (r *Rectangle) OnLoad(v *uikit.View) {
	v.SetRect(0, 0, r.w, r.h)
	size := v.Rect().Size()
	offsetImage, _ := ebiten.NewImage(size.X, size.Y, ebiten.FilterDefault)
	paint.FillRect(offsetImage, paint.Rect{X: 0, Y: 0, W: size.X, H: size.Y}, r.color)
	r.offsetImage = offsetImage
}

func (r *Rectangle) OnDraw(v *uikit.View, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(v.Rect().Min.X), float64(v.Rect().Min.Y))
	screen.DrawImage(r.offsetImage, op)
}
