package sprite

import (
	// import for side effect
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

type frame struct {
	width  int
	height int
	index  int
	length int
}

type position struct {
	x float64
	y float64
}

type size struct {
	w int
	h int
}

// Sprite manage part of image for certain size
type Sprite struct {
	image    *ebiten.Image
	size     size
	frame    frame
	position position
}

// New create the Sprite struct
func New(image *ebiten.Image, frameNum int) *Sprite {
	sprite := &Sprite{}
	sprite.image = image
	sprite.size.w, sprite.size.h = image.Size()
	sprite.frame.length = frameNum
	sprite.frame.index = 0
	sprite.frame.width = sprite.size.w / frameNum
	sprite.frame.height = sprite.size.h

	return sprite
}

// Size returns frame width and height of this Sprite
func (s *Sprite) Size() (int, int) {
	return s.frame.width, s.frame.height
}

// SetPosition set the position of this Sprite
func (s *Sprite) SetPosition(x, y float64) {
	s.position.x = x
	s.position.y = y
}

// Draw draws this sprite
func (s *Sprite) Draw(screen *ebiten.Image) {
	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)
	screen.DrawImage(s.image, op)
}
