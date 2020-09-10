package field

import (
	"image/color"

	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

const (
	fieldWidth  = 1000
	fieldHeight = 1000
)

// Field represents the game field
type Field struct {
	x             float64
	y             float64
	width         float64
	height        float64
	boundaryImage *ebiten.Image
	windowWidth   float64
	windowHeight  float64
}

// NewField creates new field
func NewField(windowWidth, windowHeight int) *Field {
	f := &Field{}
	f.x = fieldWidth / 2
	f.y = fieldHeight / 2
	f.width = fieldWidth
	f.height = fieldHeight
	f.windowWidth = float64(windowWidth)
	f.windowHeight = float64(windowHeight)

	borderColor := color.RGBA{0xff, 0, 0, 0x50}
	offsetImage, _ := ebiten.NewImage(int(f.width), int(f.height), ebiten.FilterDefault)
	paint.DrawRect(offsetImage, paint.Rect{X: 0, Y: 0, W: int(f.width), H: int(f.height)}, borderColor, 1)
	f.boundaryImage = offsetImage

	return f
}

// Draw draws the field
func (f *Field) Draw(screen *ebiten.Image) {
	_, h := sprite.Background.Size()
	sprite.Background.SetPosition(f.windowWidth/2, f.windowHeight/2)
	sprite.Background.DrawWithScale(screen, (f.windowHeight / float64(h)))

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-shared.OffsetX, -shared.OffsetY)
	screen.DrawImage(f.boundaryImage, op)
}

// GetLeft returns left
func (f *Field) GetLeft() float64 {
	return f.x - f.width/2
}

// GetTop returns top
func (f *Field) GetTop() float64 {
	return f.y - f.height/2
}

// GetRight returns right
func (f *Field) GetRight() float64 {
	return f.x + f.width/2
}

// GetBottom returns bottom
func (f *Field) GetBottom() float64 {
	return f.y + f.height/2
}

// GetCenterX returns center x
func (f *Field) GetCenterX() float64 {
	return f.width / 2
}

// GetCenterY returns center x
func (f *Field) GetCenterY() float64 {
	return f.height / 2
}
