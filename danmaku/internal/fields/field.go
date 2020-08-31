package fields

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

const (
	fieldWidth  = 240
	fieldHeight = 260
)

// Field represents the game field
type Field struct {
	x             int
	y             int
	width         int
	height        int
	boundaryImage *ebiten.Image
}

// NewField creates new field
func NewField() *Field {
	f := &Field{}
	f.x = fieldWidth / 2
	f.y = fieldHeight / 2
	f.width = fieldWidth
	f.height = fieldHeight

	borderColor := color.RGBA{0xff, 0, 0, 0x50}
	offsetImage, _ := ebiten.NewImage(f.width, f.height, ebiten.FilterDefault)
	paint.DrawRect(offsetImage, paint.Rect{X: 0, Y: 0, W: f.width, H: f.height}, borderColor, 1)
	f.boundaryImage = offsetImage

	return f
}

// Draw draws the field
func (f *Field) Draw(screen *ebiten.Image) {
	sprite.Background.SetPosition(float64(f.x), float64(f.y))
	sprite.Background.Draw(screen)

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(f.boundaryImage, op)
}

// GetLeft returns left
func (f *Field) GetLeft() int {
	return f.x - f.width/2
}

// GetTop returns top
func (f *Field) GetTop() int {
	return f.y - f.height/2
}

// GetRight returns right
func (f *Field) GetRight() int {
	return f.x + f.width/2
}

// GetBottom returns bottom
func (f *Field) GetBottom() int {
	return f.y + f.height/2
}
