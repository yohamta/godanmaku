package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/godanmaku/danmaku/internal/paint"
)

const (
	joyStickMaxAlpha         = 0xcc
	joyStickRadius   float64 = 50
)

// Joystick represents virtual keyboard
type Joystick struct {
	panelSize      int
	keySize        int
	panelNum       int
	color          color.RGBA
	offsetImage    *ebiten.Image
	animateAlpha   int
	center         position
	isReadingTouch bool
	touchID        int
	isFirstTouch   bool
}

// NewJoystick returns Joystick
func NewJoystick() *Joystick {
	joystick := &Joystick{}

	// Prepare an offset image for Joystick
	joystick.keySize = 20
	joystick.panelNum = 5
	joystick.panelSize = joystick.keySize * joystick.panelNum

	// color setting
	joystick.color = color.RGBA{0, 0xff, 0, 0xff}
	joystick.animateAlpha = -3

	// position
	scw := ScreenWidth
	sch := ScreenHeight
	joystick.center.x = scw/2 - joystick.panelSize/2
	joystick.center.y = sch - joystick.panelSize/2 - 40
	joystick.isFirstTouch = true

	joystick.createOffsetImage()

	return joystick
}

func (joystick *Joystick) createOffsetImage() {
	panelSize := joystick.panelSize
	offsetImage, _ := ebiten.NewImage(panelSize, panelSize, ebiten.FilterDefault)

	panelNum := joystick.panelNum
	keySize := joystick.keySize
	color := joystick.color
	for i := 0; i < panelNum; i++ {
		for j := 0; j < panelNum; j++ {
			x := i * keySize
			y := j * keySize
			paint.DrawRect(offsetImage, paint.Rect{X: x, Y: y, W: keySize, H: keySize}, color, 1)
		}
	}

	joystick.offsetImage = offsetImage
}

// Update updates the state of the panel
func (joystick *Joystick) Update() {
	joystick.updateColor()
}

func (joystick *Joystick) updateColor() {
	// animate the panel color
	clr := joystick.color
	a := clr.A
	a = uint8(math.Min(math.Max(float64(a)+float64(joystick.animateAlpha), 0), joyStickMaxAlpha))
	clr.A = a
	if a == joyStickMaxAlpha || a == 0 {
		joystick.animateAlpha *= -1
	}
	joystick.color = clr
}

// StartReadingTouch sets the touchID of this joystick read
func (joystick *Joystick) StartReadingTouch(touchID int) {
	joystick.touchID = touchID
	x, y := ebiten.TouchPosition(touchID)
	joystick.center.x = x
	joystick.center.y = y
	joystick.isReadingTouch = true
	joystick.isFirstTouch = false
}

// ReadInput returns current input
func (joystick *Joystick) ReadInput() (float64, float64) {
	if joystick.isReadingTouch == false {
		return 0, 0
	}
	x, y := ebiten.TouchPosition(joystick.touchID)
	dx := x - joystick.center.x
	dy := y - joystick.center.y
	horizontal := math.Min(math.Max(float64(dx)/joyStickRadius, -1.0), 1.0)
	vertical := math.Min(math.Max(float64(dy)/joyStickRadius, -1.0), 1.0)
	return horizontal, vertical
}

// IsReadingTouch returns if this joystick is currently on touch
func (joystick *Joystick) IsReadingTouch() bool {
	return joystick.isReadingTouch
}

// GetTouchID returns touch ID
func (joystick *Joystick) GetTouchID() int {
	return joystick.touchID
}

// EndReadingTouch ends touch
func (joystick *Joystick) EndReadingTouch() {
	joystick.isReadingTouch = false
}

// Draw draws joystick
func (joystick *Joystick) Draw(screen *ebiten.Image) {
	if joystick.IsReadingTouch() == false && joystick.isFirstTouch == false {
		return
	}
	op := &ebiten.DrawImageOptions{}

	// set position
	panelSize := joystick.panelSize
	x := joystick.center.x
	y := joystick.center.y
	op.GeoM.Translate(float64(x-panelSize/2), float64(y-panelSize/2))

	// set color
	c := joystick.color
	r := float64(c.R) / 0xff
	g := float64(c.G) / 0xff
	b := float64(c.B) / 0xff
	a := float64(c.A) / 0xff * -1
	op.ColorM.Translate(r, g, b, a)

	screen.DrawImage(joystick.offsetImage, op)
}
