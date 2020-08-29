package button

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
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
	isTouching bool
	touchID    int
}

// New returns new button
func New(x, y, w, h int) *Button {
	button := &Button{}
	button.position.x = x
	button.position.y = y
	button.size.w = w
	button.size.h = h
	button.isTouching = false
	return button
}

// Draw draws button to the screen
func (button *Button) Draw(screen *ebiten.Image) {
	x := button.position.x
	y := button.position.y
	w := button.size.w
	h := button.size.h
	paint.DrawRect(screen, paint.Rect{X: x, Y: y, W: w, H: h}, color.RGBA{0, 0xff, 0, 0xff}, 1)
}

// HandleTouch handle the touch and returns if it is handled or not
func (button *Button) HandleTouch(touchID, x, y int) bool {
	if button.isInside(x, y) {
		button.isTouching = true
		button.touchID = touchID
		return true
	}
	button.isTouching = false
	return false
}

func (button *Button) isInside(x, y int) bool {
	bx := button.position.x
	by := button.position.y
	bw := button.size.w
	bh := button.size.h
	return bx <= x && x <= bx+bw && by <= y && y <= by+bh
}
