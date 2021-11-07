package sprite

import (
	"fmt"

	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/texture"
)

var spriteCache = map[string]*ganim8.Sprite{}

type Sprite2Data struct {
	Id  string
	Tex *texture.Texture
}

func CacheSprite(data []*Sprite2Data) {
	for _, v := range data {
		frames := v.Tex.Grid.GetFrames()
		if len(frames) == 0 {
			panic(fmt.Sprintf("no frames, %s", v.Id))
		}
		spr := ganim8.NewSprite(v.Tex.Image, frames)
		spriteCache[v.Id] = spr
	}
}

func Get(id string) *ganim8.Sprite {
	return spriteCache[id]
}
