package input

import (
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
)

const (
	minAlphaTranslate = 0xcc
)

// Panel represents virtual keyboard
type Panel struct {
	panelSize    int
	keySize      int
	panelNum     int
	color        color.RGBA
	offsetImage  *ebiten.Image
	animateAlpha int
	lastTickTime time.Time
}

// NewPanel returns Panel
func NewPanel() *Panel {
	p := &Panel{}

	// Prepare an offset image for Panel
	p.keySize = 20
	p.panelNum = 5
	p.panelSize = p.keySize * p.panelNum

	// color setting
	p.color = color.RGBA{0, 0xff, 0, 0xff}
	p.animateAlpha = -3
	p.lastTickTime = time.Now()

	p.preparePanel()

	return p
}

func (p *Panel) preparePanel() {
	offsetImage, _ := ebiten.NewImage(p.panelSize, p.panelSize, ebiten.FilterDefault)

	// draw keys
	for i := 0; i < p.panelNum; i++ {
		for j := 0; j < p.panelNum; j++ {
			x := i * p.keySize
			y := j * p.keySize
			paint.DrawRect(offsetImage, paint.Rect{X: x, Y: y, W: p.keySize, H: p.keySize}, p.color, 1)
		}
	}

	p.offsetImage = offsetImage
}

// UpdatePanel updates the state of the panel
func (p *Panel) UpdatePanel() {
	p.updateColor()
}

func (p *Panel) updateColor() {
	// animate the panel color
	clr := p.color
	a := clr.A
	a = uint8(math.Min(math.Max(float64(a)+float64(p.animateAlpha), 0), minAlphaTranslate))
	clr.A = a
	if a == minAlphaTranslate || a == 0 {
		p.animateAlpha *= -1
	}
	p.color = clr

}

// DrawPanel draws panel
func (p *Panel) DrawPanel(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}

	// set position
	op.GeoM.Translate(float64(x-p.panelSize/2), float64(y-p.panelSize/2))

	// set color
	c := p.color
	r := float64(c.R) / 0xff
	g := float64(c.G) / 0xff
	b := float64(c.B) / 0xff
	a := float64(c.A) / 0xff * -1
	op.ColorM.Translate(r, g, b, a)

	screen.DrawImage(p.offsetImage, op)
}
