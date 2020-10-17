package shooting

import (
	"image"
	"image/color"
	"math"

	"github.com/yotahamada/furex"
	"github.com/yotahamada/godanmaku/danmaku/internal/paint"

	"github.com/hajimehoshi/ebiten"
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
}

func NewFireButton() *FireButton {
	fb := new(FireButton)
	fb.alpha = fireButtonMaxAlpha
	fb.animateAlpha = -1
	return fb
}

func (fb *FireButton) GetSize() image.Point {
	return image.Pt(fireButtonWidth, fireButtonHeight)
}

func (fb *FireButton) GetPosition() image.Point {
	x := screenSize.X/2 + screenSize.X/4 - fireButtonWidth/2
	y := screenSize.Y - fireButtonHeight - 40
	return image.Pt(x, y)
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

func (fb *FireButton) OnPressButton() {
	fb.isPressing = true
}

func (fb *FireButton) OnReleaseButton() {
	fb.isPressing = false
}

func (fb *FireButton) Draw(screen *ebiten.Image, frame image.Rectangle) {
	// TODO: performance improvement
	if fb.isPressing {
		furex.FillRect(screen, frame, color.RGBA{0xff, 0, 0, 0x50})
	} else {
		furex.FillRect(screen, frame, color.RGBA{0, 0xff, 0, 0x50})
	}
	furex.DrawRect(screen, frame, color.RGBA{0xcc, 0xcc, 0, 0x60}, 1)
	paint.DrawText(screen, "Attack", frame.Min.X+(frame.Max.X-frame.Min.X)/2-34, frame.Min.Y+(frame.Max.Y-frame.Min.Y)/2+8,
		color.White, paint.FontSizeXLarge)
}
