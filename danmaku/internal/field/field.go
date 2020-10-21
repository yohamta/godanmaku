package field

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/yohamta/godanmaku/danmaku/internal/util"

	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
)

const (
	fieldWidth  = 500
	fieldHeight = 500
)

type Field struct {
	x             float64
	y             float64
	width         float64
	height        float64
	boundaryImage *ebiten.Image
}

func NewField() *Field {
	f := &Field{}
	f.x = fieldWidth / 2
	f.y = fieldHeight / 2
	f.width = fieldWidth
	f.height = fieldHeight

	borderColor := color.RGBA{0xff, 0, 0, 0x50}
	offsetImage, _ := ebiten.NewImage(int(f.width), int(f.height), ebiten.FilterDefault)
	paint.DrawRect(offsetImage, image.Rect(0, 0, int(f.width), int(f.height)), borderColor, 1)
	f.boundaryImage = offsetImage

	return f
}

func (f *Field) GetRandamPosition(centerX, centerY, radius float64) (float64, float64) {
	rad := util.DegToRad(int(rand.Float64() * 360))
	x := math.Max(math.Min(math.Cos(rad)*rand.Float64()*radius+centerX, f.width), 0)
	y := math.Max(math.Min(math.Sin(rad)*rand.Float64()*radius+centerY, f.height), 0)
	return x, y
}

func (f *Field) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-shared.OffsetX, -shared.OffsetY)
	screen.DrawImage(f.boundaryImage, op)
}

func (f *Field) GetLeft() float64 {
	return f.x - f.width/2
}

func (f *Field) GetTop() float64 {
	return f.y - f.height/2
}

func (f *Field) GetRight() float64 {
	return f.x + f.width/2
}

func (f *Field) GetBottom() float64 {
	return f.y + f.height/2
}

func (f *Field) GetCenterX() float64 {
	return f.width / 2
}

func (f *Field) GetCenterY() float64 {
	return f.height / 2
}
