package sprite

import (
	"bytes"
	"image"

	// import for side effect
	_ "image/png"

	"github.com/hajimehoshi/ebiten"

	"github.com/yohamta/godanmaku/danmaku/internal/resources/images"
)

var (
	Background,
	Player,
	PlayerBullet,
	Enemy1,
	Enemy2 *Sprite
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

// Sprite manage part of image for certain size
type Sprite struct {
	image     *ebiten.Image
	subImages []*ebiten.Image
	size      size
	frame     frame
	position  position
	index     int
	length    int
}

// NewSprite create the Sprite struct
func NewSprite(img *image.Image, frameNum int) *Sprite {
	originalImage, _ := ebiten.NewImageFromImage(*img, ebiten.FilterDefault)

	sprite := &Sprite{}
	sprite.image = originalImage
	sprite.size.w, sprite.size.h = originalImage.Size()
	sprite.length = frameNum
	sprite.index = 0
	sprite.frame.w = sprite.size.w / frameNum
	sprite.frame.h = sprite.size.h

	subImages := []*ebiten.Image{}
	for i := 0; i < frameNum; i++ {
		x := sprite.frame.w * i
		y := 0
		rect := image.Rect(x, y, x+sprite.frame.w, y+sprite.frame.h)
		sub := originalImage.SubImage(rect)
		ebitenImage, _ := ebiten.NewImageFromImage(sub, ebiten.FilterDefault)
		subImages = append(subImages, ebitenImage)
	}
	sprite.subImages = subImages

	return sprite
}

// Size returns frame width and height of the Sprite
func (s *Sprite) Size() (int, int) {
	return s.frame.w, s.frame.h
}

// SetPosition sets the position of the Sprite
func (s *Sprite) SetPosition(x, y float64) {
	s.position.x = x
	s.position.y = y
}

// SetIndex sets the current frame of the Sprite
func (s *Sprite) SetIndex(index int) {
	s.index = index
}

// Index returns the current index of the Sprite
func (s *Sprite) Index() int {
	return s.index
}

// Length returns the length of the Sprite
func (s *Sprite) Length() int {
	return s.index
}

// Draw draws this sprite
func (s *Sprite) Draw(screen *ebiten.Image) {
	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)

	screen.DrawImage(s.subImages[s.index], op)
}

// LoadSprites loads sprites
func LoadSprites() {
	Player = createSprite(&images.PLAYER, 8)
	Background = createSprite(&images.UTYUU_BG, 1)
	PlayerBullet = createSprite(&images.SHOT2, 8)
	Enemy1 = createSprite(&images.ENEMY1, 8)
	Enemy2 = createSprite(&images.ENEMY2, 1)
}

func createSprite(rawImage *[]byte, frameNum int) *Sprite {
	img, _, _ := image.Decode(bytes.NewReader(*rawImage))
	return NewSprite(&img, 8)
}
