package paint

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/godanmaku/danmaku/internal/resources/fonts"
	"golang.org/x/image/font"

	// "github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/text"
)

// FontSize represents font size
type FontSize int

const (
	// FontSizeXLarge represents xlarge font
	FontSizeXLarge FontSize = iota
)

var (
	pixelMPlusRegular font.Face
)

// LoadFonts loads fonts
func LoadFonts() {
	tt, err := truetype.Parse(fonts.PIXELMPLUS12REGULAR)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	pixelMPlusRegular = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

// DrawText draws text
func DrawText(target *ebiten.Image, txt string, x, y int, clr color.Color, fontSize FontSize) {
	text.Draw(target, txt, pixelMPlusRegular, x, y, clr)
}
