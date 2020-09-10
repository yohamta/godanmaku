package ui

import (
	"github.com/hajimehoshi/ebiten"
)

// View represents view to draw on screen
type View interface {
	Update()
	Draw(screen *ebiten.Image)
	GetPosition() (x, y int)
	GetSize() (width, height int)
}

var (
	screenWidth, screenHeight int
	viewStack                 *stack
)

type position struct {
	x, y int
}

type size struct {
	w int
	h int
}

// GetScreenWidth returns width of the screen
func GetScreenWidth() int {
	return screenWidth
}

// GetScreenHeight returns height of the screen
func GetScreenHeight() int {
	return screenHeight
}

// GetScreenSize returns width of the screen
func GetScreenSize() (int, int) {
	return screenWidth, screenHeight
}

// SetRootView set the root view
func SetRootView(v View) {
	viewStack = newStack()
	viewStack.push(v)
}

// SetScreenSize returns width of the screen
func SetScreenSize(width, height int) {
	screenWidth = width
	screenHeight = height
}

// GetCenterOfScreen returns center position of the screen
func GetCenterOfScreen() (x, y int) {
	return screenWidth / 2, screenHeight / 2
}

// Update updates the screen
func Update() {
	for ite := viewStack.getIterator(); ite.hasNext(); {
		v := ite.next()
		if v != nil {
			v.Update()
		}
	}
}

// Draw draws the screen
func Draw(screen *ebiten.Image) {
	if viewStack.size() == 0 {
		return
	}
	for ite := viewStack.getIterator(); ite.hasNext(); {
		v := ite.next()
		if v != nil {
			v.Draw(screen)
		}
	}
}
