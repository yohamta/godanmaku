package sprite

import (
	"image"

	_ "image/png"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type frame struct {
	w int
	h int
}

type position struct {
	x float64
	y float64
}

type size struct {
	w int
	h int
}

type Sprite struct {
	image     *ebiten.Image
	subImages []*ebiten.Image
	size      size
	frame     frame
	position  position
	index     int
	length    int
}

func NewSprite(img *image.Image, columns int, rows int) *Sprite {
	originalImage := ebiten.NewImageFromImage(*img)

	sprite := &Sprite{}
	sprite.image = originalImage
	sprite.size.w, sprite.size.h = originalImage.Size()
	sprite.length = columns * rows
	sprite.index = 0

	sprite.frame.w = sprite.size.w / columns
	sprite.frame.h = sprite.size.h / rows

	subImages := []*ebiten.Image{}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			y := sprite.frame.h * i
			x := sprite.frame.w * j
			rect := image.Rect(x, y, x+sprite.frame.w, y+sprite.frame.h)
			sub := originalImage.SubImage(rect)
			ebitenImage := ebiten.NewImageFromImage(sub)
			subImages = append(subImages, ebitenImage)
		}
	}
	sprite.subImages = subImages

	return sprite
}

func (s *Sprite) Size() (int, int) {
	return s.frame.w, s.frame.h
}

func (s *Sprite) GetWidth() int {
	return s.frame.w
}

func (s *Sprite) GetHeight() int {
	return s.frame.h
}

func (s *Sprite) SetPosition(x, y float64) {
	s.position.x = x
	s.position.y = y
}

func (s *Sprite) SetIndex(index int) {
	s.index = index
}

func (s *Sprite) Index() int {
	return s.index
}

func (s *Sprite) Length() int {
	return s.length
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)

	screen.DrawImage(s.subImages[s.index], op)
}

func (s *Sprite) DrawAdditive(screen *ebiten.Image, strength float64, scaleW float64, scaleH float64) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x/scaleW-float64(w)/2, y/scaleH-float64(h)/2)

	op.GeoM.Scale(scaleW, scaleH)

	op.ColorM.Translate(0, 0, 0, -1+strength)

	op.CompositeMode = ebiten.CompositeModeLighter

	screen.DrawImage(s.subImages[s.index], op)
}

func (s *Sprite) DrawWithTint(screen *ebiten.Image, r, g, b, a float64) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)

	op.ColorM.Translate(r, g, b, a)

	screen.DrawImage(s.subImages[s.index], op)
}

func (s *Sprite) DrawWithHsv(screen *ebiten.Image, hue, sat, val float64) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)

	op.ColorM.ChangeHSV(hue, sat, val)

	screen.DrawImage(s.subImages[s.index], op)
}

func (s *Sprite) DrawWithScale(screen *ebiten.Image, scale float64) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x/scale-float64(w)/2, y/scale-float64(h)/2)

	op.GeoM.Scale(scale, scale)

	screen.DrawImage(s.subImages[s.index], op)
}

func (s *Sprite) DrawWithScaleRotate(screen *ebiten.Image, scale float64, rotate float64) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Rotate(rotate)

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x/scale-float64(w)/2, y/scale-float64(h)/2)

	op.GeoM.Scale(scale, scale)

	screen.DrawImage(s.subImages[s.index], op)
}
