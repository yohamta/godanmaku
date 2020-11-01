package shooting

import "image/color"

var (
	textColorWarn        = color.RGBA{R: 0xff, G: 0x8c, B: 0xff, A: 0xff}
	textColorAchievement = color.RGBA{R: 0xa8, G: 0xff, B: 0xff, A: 0xff}
	textColorInfo        = color.RGBA{R: 0xbf, G: 0xff, B: 0x7f, A: 0xff}
)

var (
	textJp = map[string]string{
		"NEW_ENEMY_APPEAR": "敵の増援が出現!",
		"ITEM_APPEAR":      "アイテム出現!",
		"DESTROY_ENEMY":    "%s を撃破!",
	}
)

func getText(id string) string {
	return textJp[id]
}
