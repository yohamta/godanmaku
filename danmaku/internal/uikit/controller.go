package uikit

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type ViewController struct {
	frame    image.Rectangle
	rootView *View
}

func NewViewController() *ViewController {
	c := &ViewController{}

	return c
}

func (c *ViewController) SetFrame(x, y, w, h int) {
	c.frame = image.Rect(x, y, x+w, y+h)
	if c.rootView != nil {
		c.rootView.SetPosition(x, y)
		c.rootView.SetSize(w, h)
	}
}

func (c *ViewController) SetRootView(v *View) {
	c.rootView = v
	c.rootView.SetPosition(c.frame.Min.X, c.frame.Min.Y)
	c.rootView.SetSize(c.frame.Size().X, c.frame.Size().Y)
	c.rootView.Load()
}

func (c *ViewController) Update() {
	c.rootView.Update()
}

func (c *ViewController) Layout() {
	c.rootView.Layout()
}

func (c *ViewController) Draw(screen *ebiten.Image) {
	c.rootView.Draw(screen)
}
