package button

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"

	"github.com/yohamta/godanmaku/danmaku/internal/paint"
)

type position struct {
	x int
	y int
}

type size struct {
	w int
	h int
}

// Button represents button
type Button struct {
	position   position
	size       size
	isPressing bool
	touchID    int
}

// New returns new button
func New(x, y, w, h int) *Button {
	button := &Button{}
	button.position.x = x
	button.position.y = y
	button.size.w = w
	button.size.h = h
	button.isPressing = false
	return button
}

// Draw draws button to the screen
// This method should be overrided by other struct
func (button *Button) Draw(screen *ebiten.Image) {
	x := button.position.x
	y := button.position.y
	w := button.size.w
	h := button.size.h
	paint.DrawRect(screen, paint.Rect{X: x, Y: y, W: w, H: h}, color.RGBA{0, 0xff, 0, 0xff}, 1)
}

// Update updates the state of the button
// This method should be overrided by expansion struct
func (button *Button) Update() {
	// To be implemented by other struct
}

// HandleTouch handle the touch and returns if it is handled or not
func (button *Button) HandleTouch(touchID int) bool {
	if button.isPressing {
		return false
	}
	x, y := ebiten.TouchPosition(touchID)
	if button.isInside(x, y) {
		button.isPressing = true
		button.touchID = touchID
		return true
	}
	return false
}

// IsPressing returns if this button is pressing
func (button *Button) IsPressing() bool {
	return button.isPressing
}

// GetTouchID returns current touch ID
func (button *Button) GetTouchID() int {
	return button.touchID
}

// CheckIsTouchRelased check touch release
func (button *Button) CheckIsTouchRelased() {
	if inpututil.IsTouchJustReleased(button.touchID) {
		button.isPressing = false
	}
}

// GetPosition returns position
func (button *Button) GetPosition() (int, int) {
	return button.position.x, button.position.y
}

// GetSize returns position
func (button *Button) GetSize() (int, int) {
	return button.size.w, button.size.h
}

func (button *Button) isInside(x, y int) bool {
	bx := button.position.x
	by := button.position.y
	bw := button.size.w
	bh := button.size.h
	return bx <= x && x <= bx+bw && by <= y && y <= by+bh
}
