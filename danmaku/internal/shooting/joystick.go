package shooting

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
)

const (
	joyStickMaxAlpha         = 0xcc
	joyStickRadius   float64 = 50
)

type Joystick struct {
	panelSize      int
	keySize        int
	panelNum       int
	color          color.RGBA
	offsetImage    *ebiten.Image
	animateAlpha   int
	center         image.Point
	isReadingTouch bool
	touchID        ebiten.TouchID
	isFirstTouch   bool
	horizontal     float64
	vertical       float64
}

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
	scw := screenSize.X
	sch := screenSize.Y

	joystick.center.X = scw/2 - joystick.panelSize/2
	joystick.center.Y = sch - joystick.panelSize/2 - 40
	joystick.isFirstTouch = true

	joystick.createOffsetImage()

	return joystick
}

func (joystick *Joystick) GetSize() image.Point {
	return screenSize
}

func (joystick *Joystick) GetPosition() image.Point {
	return image.Pt(0, 0)
}

func (joystick *Joystick) HandleJustPressedTouchID(touchID ebiten.TouchID) bool {
	joystick.touchID = touchID
	x, y := ebiten.TouchPosition(touchID)
	joystick.center.X = x
	joystick.center.Y = y
	joystick.isReadingTouch = true
	joystick.isFirstTouch = false

	return true
}

func (joystick *Joystick) HandleJustReleasedTouchID(touchID int) {
	joystick.isReadingTouch = false
}

func (joystick *Joystick) createOffsetImage() {
	panelSize := joystick.panelSize
	offsetImage := ebiten.NewImage(panelSize, panelSize)

	panelNum := joystick.panelNum
	keySize := joystick.keySize
	color := joystick.color
	for i := 0; i < panelNum; i++ {
		for j := 0; j < panelNum; j++ {
			x := i * keySize
			y := j * keySize
			paint.DrawRect(offsetImage, image.Rect(x, y, x+keySize, y+keySize), color, 1)
		}
	}

	joystick.offsetImage = offsetImage
}

func (joystick *Joystick) Update() {
	joystick.updateColor()
	joystick.readInput()
}

func (joystick *Joystick) Draw(screen *ebiten.Image, frame image.Rectangle) {
	if joystick.isReadingTouch == false && joystick.isFirstTouch == false {
		return
	}
	op := &ebiten.DrawImageOptions{}

	// set position
	panelSize := joystick.panelSize
	x := joystick.center.X
	y := joystick.center.Y
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

func (joystick *Joystick) updateColor() {
	clr := joystick.color
	a := clr.A
	a = uint8(math.Min(math.Max(float64(a)+float64(joystick.animateAlpha), 0), joyStickMaxAlpha))
	clr.A = a
	if a == joyStickMaxAlpha || a == 0 {
		joystick.animateAlpha *= -1
	}
	joystick.color = clr
}

func (joystick *Joystick) readInput() {
	if joystick.isReadingTouch {
		x, y := ebiten.TouchPosition(joystick.touchID)
		dx := x - joystick.center.X
		dy := y - joystick.center.Y
		joystick.horizontal = math.Min(math.Max(float64(dx)/joyStickRadius, -1.0), 1.0)
		joystick.vertical = math.Min(math.Max(float64(dy)/joyStickRadius, -1.0), 1.0)
	}
}
