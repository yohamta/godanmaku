package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Button represents button
type Button struct {
	position   position
	size       size
	isPressing bool
	touchID    int
}

// NewButton returns new button
func NewButton(x, y, w, h int) *Button {
	button := &Button{}
	button.position.x = x
	button.position.y = y
	button.size.w = w
	button.size.h = h
	button.isPressing = false
	return button
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
