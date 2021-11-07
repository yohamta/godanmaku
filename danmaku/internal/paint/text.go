package paint

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/resources/fonts"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2/text"
)

type FontSize int

const (
	FontSizeXLarge FontSize = iota
	FontSizeMedium
	FontSizeSmall
)

type HAlign int

const (
	HAlignLeft HAlign = iota
	HAlignCenter
)

var (
	fontMap = map[FontSize]font.Face{}
)

func LoadFonts() {
	tt, err := truetype.Parse(fonts.SOUKOUMINCHO)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	fontMap[FontSizeXLarge] = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	fontMap[FontSizeMedium] = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	fontMap[FontSizeSmall] = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

type DrawTextOptions struct {
	Color    color.Color
	Width    int
	HAligh   HAlign
	FontSize FontSize
}

func DrawText(target *ebiten.Image, txt string, x, y int, clr color.Color, fontSize FontSize) {
	text.Draw(target, txt, fontMap[fontSize], x, y, clr)
}

func DrawTextWithOptions(target *ebiten.Image, txt string, x, y int, opts DrawTextOptions) {
	f := fontMap[opts.FontSize]
	c := opts.Color
	r := text.BoundString(f, txt)
	x2 := 0
	y2 := 0
	if opts.HAligh == HAlignLeft {
		x2 = -r.Min.X + x
		y2 = -r.Min.Y + y
	} else if opts.HAligh == HAlignCenter {
		x2 = x - (r.Max.X-r.Min.X)/2
		y2 = y + (r.Max.Y-r.Min.Y)/2
	}
	text.Draw(target, txt, f, x2, y2, c)
}
