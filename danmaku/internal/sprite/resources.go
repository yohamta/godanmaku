package sprite

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"math/rand"

	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/resources/images"
	"github.com/yohamta/godanmaku/danmaku/internal/texture"
)

var EnemyShots = []string{
	"ESHOT10_1",
	"ESHOT10_2",
	"ESHOT10_3",
	"ESHOT10_4",
	"ESHOT10_5",
	"ESHOT10_6",
}

var textureData = []*texture.TextureData{
	{"P_ROBO1", "main", decodeImage(&images.P_ROBO1), 16, 16},
	{"BACKGROUND", "background", decodeImage(&images.SPACE5), 240, 260},
	{"PSHOT_1", "main", decodeImage(&images.SHOT1), 10, 10},
	{"E_ROBO1", "main", decodeImage(&images.E_ROBO1), 24, 24},
	{"HIT", "main", decodeImage(&images.HIT_SMALL), 28, 28},
	{"EXPLODE_SMALL", "main", decodeImage(&images.EXPLODE_SMALL), 32, 32},
	{"JUMP", "main", decodeImage(&images.JUMP), 64, 64},
	{"BATTLE_RESULT", "ui", decodeImage(&images.SYOUHAI), 56, 31},
	{"TRAIL", "main", decodeImage(&images.KISEKI), 16, 16},
	{"BLUE_LASER", "main", decodeImage(&images.RASER1), 16, 16},
	{"BLUE_LASER_LONG", "main", decodeImage(&images.RASERLONG1), 16, 16},
	{"TRAIL", "main", decodeImage(&images.BACKFIRE), 32, 32},
	{"FUNNEL", "main", decodeImage(&images.PBIT_1), 10, 10},
	{"ITEM_P", "main", decodeImage(&images.ITEM_P), 16, 17},
	{"ITEM_LIFE", "main", decodeImage(&images.ITEM_LIFE), 16, 10},
	{"GRAZE", "main", decodeImage(&images.KASRUI), 16, 16},
	{"ESHOT10_1", "main", decodeImage(&images.ESHOT10_1), 10, 10},
	{"ESHOT10_2", "main", decodeImage(&images.ESHOT10_2), 10, 10},
	{"ESHOT10_3", "main", decodeImage(&images.ESHOT10_3), 10, 10},
	{"ESHOT10_4", "main", decodeImage(&images.ESHOT10_4), 10, 10},
	{"ESHOT10_5", "main", decodeImage(&images.ESHOT10_5), 10, 10},
	{"ESHOT10_6", "main", decodeImage(&images.ESHOT10_6), 10, 10},
}

func LoadSprites() {
	texture.LoadTextures(textureData)
	var spriteData = []*Sprite2Data{}
	for _, v := range textureData {
		texture := texture.GetTexure(v.Id)
		spriteData = append(spriteData, &Sprite2Data{
			Id:  v.Id,
			Tex: texture,
		})
		if texture == nil {
			panic(fmt.Sprintf("texture is nil, %s", v.Id))
		}
	}
	CacheSprite(spriteData)
}

func RandomEnemyShot() *ganim8.Sprite {
	i := rand.Intn(len(EnemyShots))
	return Get(EnemyShots[i])
}

func decodeImage(rawImage *[]byte) *image.Image {
	img, _, _ := image.Decode(bytes.NewReader(*rawImage))
	return &img
}
