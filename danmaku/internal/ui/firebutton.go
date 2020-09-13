package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
)

const (
	fireButtonWidth    = 80
	fireButtonHeight   = 80
	fireButtonMaxAlpha = 0x30
)

// FireButton represents Button
type FireButton struct {
	Button
	onImage      *ebiten.Image
	offImage     *ebiten.Image
	alpha        uint8
	animateAlpha int
}

// NewFireButton returns new FireButton
func NewFireButton() *FireButton {
	scw, sch := GetScreenSize()
	x := scw/2 + scw/4 - fireButtonWidth/2
	y := sch - fireButtonHeight - 40
	baseButton := (NewButton(x, y, fireButtonWidth, fireButtonHeight))
	fButton := &FireButton{Button: *baseButton}

	// visual setting
	fButton.alpha = fireButtonMaxAlpha
	fButton.animateAlpha = -1
	fButton.offImage = initOffsetImage(
		color.RGBA{0, 0xff, 0, 0x50},
		color.RGBA{0, 0xcc, 0, 0x60})
	fButton.onImage = initOffsetImage(
		color.RGBA{0xff, 0xff, 0, 0x50},
		color.RGBA{0xcc, 0xcc, 0, 0x60})

	return fButton
}

func initOffsetImage(bgClr, bdrClr color.RGBA) *ebiten.Image {
	width := fireButtonWidth
	height := fireButtonHeight
	offImage, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	paint.FillRect(offImage, paint.Rect{X: 0, Y: 0, W: width, H: height}, bgClr)
	paint.DrawRect(offImage, paint.Rect{X: 0, Y: 0, W: width, H: height}, bdrClr, 1)
	paint.DrawText(offImage, "Attack", width/2-34, height/2+8, color.White, paint.FontSizeXLarge)
	return offImage
}

func (fButton *FireButton) updateColor() {
	maxAlpha := fireButtonMaxAlpha
	a := fButton.alpha
	a = uint8(math.Min(math.Max(float64(a)+float64(fButton.animateAlpha), 0), float64(maxAlpha)))
	if a == uint8(maxAlpha) || a == 0 {
		fButton.animateAlpha *= -1
	}
	fButton.alpha = a
}

// Update updates the button
func (fButton *FireButton) Update() {
	fButton.updateColor()
}

// Draw draws the button
func (fButton *FireButton) Draw(screen *ebiten.Image) {
	x, y := fButton.GetPosition()

	op := &ebiten.DrawImageOptions{}

	// set position
	op.GeoM.Translate(float64(x), float64(y))

	// set color
	a := -1 * float64(fButton.alpha) / float64(0xff)
	op.ColorM.Translate(0, 0, 0, a)

	if fButton.IsPressing() {
		screen.DrawImage(fButton.onImage, op)
	} else {
		screen.DrawImage(fButton.offImage, op)
	}
}
