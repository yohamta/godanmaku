package texture

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/pckr"
)

type TextureData struct {
	Id    string
	Group string
	Image *image.Image
	W, H  int
}

type Texture struct {
	Id    string
	Image *ebiten.Image
	Grid  *ganim8.Grid
}

var packedImages = map[string]*pckr.Packer{}
var cachedGrids = map[string]*Texture{}

func LoadTextures(data []*TextureData) {
	dict := map[string][]*TextureData{}
	for _, v := range data {
		k := v.Group
		if _, ok := dict[k]; !ok {
			dict[k] = []*TextureData{}
		}
		dict[k] = append(dict[k], v)
	}
	for k, v := range dict {
		packer := _pack(v)
		packedImages[k] = packer
		for _, tex := range v {
			r := packer.Location(tex.Id)
			fmt.Printf("Load sprite %s\n", tex.Id)
			cachedGrids[tex.Id] = &Texture{
				tex.Id,
				packer.Image(),
				ganim8.NewGrid(tex.W, tex.H, r.Dx(), r.Dy(), r.Min.X, r.Min.Y),
			}
		}
	}
}

func GetTexure(id string) *Texture {
	return cachedGrids[id]
}

func _pack(imageMaster []*TextureData) *pckr.Packer {
	packer := pckr.NewPacker(2048, 2048)
	for _, v := range imageMaster {
		img := ebiten.NewImageFromImage(*v.Image)
		w, h := img.Size()
		packer.Add(v.Id, img, 0, 0, w, h)
	}
	packer.Pack()
	return packer
}

func decodeImage(rawImage *[]byte) *image.Image {
	img, _, _ := image.Decode(bytes.NewReader(*rawImage))
	return &img
}
