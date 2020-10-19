package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
)

// HealthBar represents health bar ui
type HealthBar struct {
	barBorder  *ebiten.Image
	barInside  *ebiten.Image
	barInside2 *ebiten.Image
}

const (
	barWidth  = 24.
	barHeight = 3.
)

// NewHealthBar create new HealthBar
func NewHealthBar() *HealthBar {
	b := &HealthBar{}

	b.createOffsetImage()

	return b
}

func (b *HealthBar) createOffsetImage() {
	img1, _ := ebiten.NewImage(barWidth, barHeight, ebiten.FilterDefault)
	borderColor := color.RGBA{0x60, 0x60, 0x60, 0xff}
	paint.DrawRect(img1, paint.Rect{X: 0, Y: 0, W: barWidth, H: barHeight}, borderColor, 1)
	b.barBorder = img1

	img2, _ := ebiten.NewImage(barWidth, barHeight, ebiten.FilterDefault)
	c2 := color.RGBA{0x66, 0xff, 0x66, 0xff}
	paint.FillRect(img2, paint.Rect{X: 0, Y: 0, W: 1, H: barHeight}, c2)
	b.barInside = img2

	img3, _ := ebiten.NewImage(barWidth, barHeight, ebiten.FilterDefault)
	c3 := color.RGBA{0xff, 0x66, 0x66, 0xff}
	paint.FillRect(img3, paint.Rect{X: 0, Y: 0, W: 1, H: barHeight}, c3)
	b.barInside2 = img3
}

// Draw draws health bar
func (b *HealthBar) Draw(x, y, ratio float64, screen *ebiten.Image) {
	// bar
	scale := barWidth * ratio
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((x-barWidth/2)/scale, y-barHeight/2)
	op.GeoM.Scale(scale, 1)
	if ratio < 0.2 {
		screen.DrawImage(b.barInside2, op)
	} else {
		screen.DrawImage(b.barInside, op)
	}

	// border
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(x-barWidth/2, y-barHeight/2)
	screen.DrawImage(b.barBorder, op2)

}
