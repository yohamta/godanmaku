package paint

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Rect represents an area of image
type Rect struct {
	X, Y, W, H int
}

// FillRect fills an area of the image
func FillRect(target *ebiten.Image, r Rect, clr color.Color) {
	img, _ := ebiten.NewImage(r.W, r.H, ebiten.FilterDefault)
	img.Fill(clr)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.X), float64(r.Y))

	target.DrawImage(img, op)
}

// DrawRect draws rect
func DrawRect(target *ebiten.Image, r Rect, clr color.Color, width int) {
	for i := r.X; i < r.X+r.W; i++ {
		for j := 0; j < width; j++ {
			target.Set(i, r.Y+j, clr)
			target.Set(i, r.Y+r.H-j-1, clr)
		}
	}

	for i := r.Y; i < r.Y+r.H; i++ {
		for j := 0; j < width; j++ {
			target.Set(r.X+j, i, clr)
			target.Set(r.X+r.W-j-1, i, clr)
		}
	}
}
