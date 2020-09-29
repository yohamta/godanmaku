package sprite

import (
	"bytes"
	"image"
	"math/rand"

	"github.com/yotahamada/godanmaku/danmaku/internal/resources/images"
)

var (
	Background    *Sprite
	Player        *Sprite
	Enemy1        *Sprite
	Enemy2        *Sprite
	Hit           *Sprite
	Explosion     *Sprite
	Jump          *Sprite
	Result        *Sprite
	Locus         *Sprite
	Nova          *Sprite
	PlayerBullet  *Sprite
	EnemyShots    []*Sprite
	BlueLaser     *Sprite
	Sparkle       *Sprite
	Backfire      *Sprite
	BlueLaserLong *Sprite
)

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
	BlueLaser = createSprite(&images.RASER1, 6, 4)
	BlueLaserLong = createSprite(&images.RASERLONG1, 6, 4)
	Sparkle = createSprite(&images.SPARKLE, 8, 8)
	Backfire = createSprite(&images.BACKFIRE, 8, 1)

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
