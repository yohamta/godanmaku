package sprite

import (
	"bytes"
	"image"
	"math/rand"

	// import for side effect
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/resources/images"
)

var (
	Background   *Sprite
	Player       *Sprite
	PlayerBullet *Sprite
	Enemy1       *Sprite
	Enemy2       *Sprite
	Hit          *Sprite
	Explosion    *Sprite
	Jump         *Sprite
	EnemyShots   []*Sprite
	Result       *Sprite
	Locus        *Sprite
	Nova         *Sprite
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
func NewSprite(img *image.Image, columns int, rows int) *Sprite {
	originalImage, _ := ebiten.NewImageFromImage(*img, ebiten.FilterDefault)

	sprite := &Sprite{}
	sprite.image = originalImage
	sprite.size.w, sprite.size.h = originalImage.Size()
	sprite.length = columns
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
			ebitenImage, _ := ebiten.NewImageFromImage(sub, ebiten.FilterDefault)
			subImages = append(subImages, ebitenImage)
		}
	}
	sprite.subImages = subImages

	return sprite
}

// Size returns frame width and height of the Sprite
func (s *Sprite) Size() (int, int) {
	return s.frame.w, s.frame.h
}

// GetWidth returns frame width and height of the Sprite
func (s *Sprite) GetWidth() int {
	return s.frame.w
}

// GetHeight returns frame width and height of the Sprite
func (s *Sprite) GetHeight() int {
	return s.frame.h
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
	return s.length
}

// Draw draws this sprite
func (s *Sprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)

	screen.DrawImage(s.subImages[s.index], op)
}

// DrawAdditive draws additive image
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

// DrawWithTint draws this sprite
func (s *Sprite) DrawWithTint(screen *ebiten.Image, r, g, b, a float64) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)

	op.ColorM.Translate(r, g, b, a)

	screen.DrawImage(s.subImages[s.index], op)
}

// DrawWithHsv draws this sprite
func (s *Sprite) DrawWithHsv(screen *ebiten.Image, hue, sat, val float64) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)

	op.ColorM.ChangeHSV(hue, sat, val)

	screen.DrawImage(s.subImages[s.index], op)
}

// DrawWithScale draws this sprite
func (s *Sprite) DrawWithScale(screen *ebiten.Image, scale float64) {
	op := &ebiten.DrawImageOptions{}

	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op.GeoM.Translate(x/scale-float64(w)/2, y/scale-float64(h)/2)

	op.GeoM.Scale(scale, scale)

	screen.DrawImage(s.subImages[s.index], op)
}

// DrawWithScaleRotate draws this sprite
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

// LoadSprites loads sprites
func LoadSprites() {
	Player = createSprite(&images.P_ROBO1, 8, 1)
	Background = createSprite(&images.SPACE5, 1, 1)
	PlayerBullet = createSprite(&images.SHOT1, 1, 1)
	Enemy1 = createSprite(&images.E_ROBO1, 8, 1)
	Hit = createSprite(&images.HIT_SMALL, 8, 1)
	Explosion = createSprite(&images.EXPLODE_SMALL, 10, 1)
	Jump = createSprite(&images.JUMP, 5, 1)
	Result = createSprite(&images.SYOUHAI, 1, 3)
	Locus = createSprite(&images.KISEKI, 5, 1)
	Nova = createSprite(&images.NOVA, 1, 1)

	addEnemyShotSprite(createSprite(&images.ESHOT10_1, 1, 1))
	addEnemyShotSprite(createSprite(&images.ESHOT10_2, 1, 1))
	addEnemyShotSprite(createSprite(&images.ESHOT10_3, 1, 1))
	addEnemyShotSprite(createSprite(&images.ESHOT10_4, 1, 1))
	addEnemyShotSprite(createSprite(&images.ESHOT10_5, 1, 1))
	addEnemyShotSprite(createSprite(&images.ESHOT10_6, 1, 1))
}

// RandomEnemyShot returns random sprite for enemy shots
func RandomEnemyShot() *Sprite {
	return EnemyShots[int(rand.Float64()*float64(len(EnemyShots)))]
}

func createSprite(rawImage *[]byte, columns int, rows int) *Sprite {
	img, _, _ := image.Decode(bytes.NewReader(*rawImage))
	return NewSprite(&img, columns, rows)
}

func addEnemyShotSprite(s *Sprite) {
	EnemyShots = append(EnemyShots, s)
}
