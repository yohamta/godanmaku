package shooting

import (
	"image"
	"image/color"
	"math"

	"github.com/yohamta/godanmaku/danmaku/internal/paint"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	fireButtonWidth    = 80
	fireButtonHeight   = 80
	fireButtonMaxAlpha = 0x30
)

type FireButton struct {
	alpha        uint8
	animateAlpha int
	isPressing   bool
	offImage     *ebiten.Image
	onImage      *ebiten.Image
}

func NewFireButton() *FireButton {
	fb := new(FireButton)
	fb.alpha = fireButtonMaxAlpha
	fb.animateAlpha = -1

	fb.makeOffsetImages()

	return fb
}

func (fb *FireButton) Size() (int, int) {
	return fireButtonWidth, fireButtonHeight
}

func (fb *FireButton) Position() (int, int) {
	x := screenSize.X/2 + screenSize.X/4 - fireButtonWidth/2
	y := screenSize.Y - fireButtonHeight - 40
	return x, y
}

func (fb *FireButton) Update() {
	maxAlpha := fireButtonMaxAlpha
	a := fb.alpha
	a = uint8(math.Min(math.Max(float64(a)+float64(fb.animateAlpha), 0), float64(maxAlpha)))
	if a == uint8(maxAlpha) || a == 0 {
		fb.animateAlpha *= -1
	}
	fb.alpha = a
}

func (b *FireButton) HandlePress(x, y int) {
	b.isPressing = true
}

func (b *FireButton) HandleRelease(x, y int, isCancel bool) {
	b.isPressing = false
}

func (fb *FireButton) Draw(screen *ebiten.Image, frame image.Rectangle) {
	var img *ebiten.Image
	if fb.isPressing {
		img = fb.onImage
	} else {
		img = fb.offImage
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(frame.Min.X), float64(frame.Min.Y))
	screen.DrawImage(img, op)
}

func (fb *FireButton) makeOffsetImages() {
	fb.offImage = fb.makeOffImageForState(false)
	fb.onImage = fb.makeOffImageForState(true)
}

func (fb *FireButton) makeOffImageForState(isOn bool) *ebiten.Image {
	off := ebiten.NewImage(fireButtonWidth, fireButtonHeight)
	var cl color.RGBA
	frame := image.Rect(0, 0, fireButtonWidth, fireButtonHeight)
	if isOn {
		cl = color.RGBA{0xff, 0, 0, 0x50}
	} else {
		cl = color.RGBA{0, 0xff, 0, 0x50}
	}
	paint.FillRect(off, frame, cl)
	paint.DrawRect(off, frame, color.RGBA{0xcc, 0xcc, 0, 0x60}, 1)
	paint.DrawText(off, "攻撃", frame.Min.X+(frame.Max.X-frame.Min.X)/2-25, frame.Min.Y+(frame.Max.Y-frame.Min.Y)/2+8,
		color.Black, paint.FontSizeXLarge)
	return off
}
