package uikit

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

type View struct {
	loaded   bool
	bounds   image.Rectangle
	subViews []*View
	handler  ViewEventHandler
}

type ViewEventHandler interface {
	OnLoad(view *View)
	OnLayout(view *View)
	OnUpdate(view *View)
	OnDraw(view *View, screen *ebiten.Image)
}

type ViewEventHandlerFuncs struct {
	OnLoadFunc   func(view *View)
	OnLayoutFunc func(view *View)
	OnUpdateFunc func(view *View)
	OnDrawFunc   func(view *View, screen *ebiten.Image)
}

func (handler *ViewEventHandlerFuncs) OnLoad(view *View) {
	if handler.OnLoadFunc != nil {
		handler.OnLoadFunc(view)
	}
}

func (handler *ViewEventHandlerFuncs) OnLayout(view *View) {
	if handler.OnLayoutFunc != nil {
		handler.OnLayoutFunc(view)
	}
}

func (handler *ViewEventHandlerFuncs) OnUpdate(view *View) {
	if handler.OnUpdateFunc != nil {
		handler.OnUpdateFunc(view)
	}
}

func (handler *ViewEventHandlerFuncs) OnDraw(view *View, screen *ebiten.Image) {
	if handler.OnDrawFunc != nil {
		handler.OnDrawFunc(view, screen)
	}
}

func NewView(eventHandler ViewEventHandler) *View {
	v := &View{}

	v.handler = eventHandler

	return v
}

func (v *View) SetPosition(x, y int) {
	v.bounds.Add(image.Point{
		x - v.bounds.Min.X,
		y - v.bounds.Min.Y,
	})
}

func (v *View) SetSize(w, h int) {
	v.bounds.Max.X = w + v.bounds.Min.X
	v.bounds.Max.Y = h + v.bounds.Min.Y
}

func (v *View) GetPosition() image.Point {
	return v.bounds.Min
}

func (v *View) GetSize() image.Point {
	return v.bounds.Max.Sub(v.bounds.Min)
}

func (v *View) Rect() image.Rectangle {
	return v.bounds
}

func (v *View) SetRect(x0, y0, x1, y1 int) {
	v.bounds = image.Rect(x0, y0, x1, y1)
}

func (v *View) SetViewEventHandler(h ViewEventHandler) {
	v.handler = h
}

func (v *View) Layout() {
	if v.handler != nil {
		v.handler.OnLayout(v)
	}
	for i := 0; i < len(v.subViews); i++ {
		v.subViews[i].Layout()
	}
}

func (v *View) AddSubView(sub *View) {
	v.subViews = append(v.subViews, sub)
	sub.Load()
	v.Layout()
}

func (v *View) GetSubView() []*View {
	return v.subViews
}

func (v *View) Load() {
	if v.loaded == false {
		if v.handler != nil {
			v.handler.OnLoad(v)
		}
		v.loaded = true
	}
	for i := 0; i < len(v.subViews); i++ {
		v.subViews[i].Load()
	}
}

func (v *View) Update() {
	if v.handler != nil {
		v.handler.OnUpdate(v)
	}
	for i := 0; i < len(v.subViews); i++ {
		v.subViews[i].Update()
	}
}

func (v *View) Draw(screen *ebiten.Image) {
	if v.handler != nil {
		v.handler.OnDraw(v, screen)
	}
	for i := 0; i < len(v.subViews); i++ {
		v.subViews[i].Draw(screen)
	}
}
