package collision

import (
	"fmt"
)

// Collider represents collider
type Collider interface {
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
	GetCollisionBox() []*Box
}

// Box represents collision box
type Box struct {
	x, y, w, h float64
}

var (
	boxData = map[string]([]*Box){
		"WEAPON_NORMAL_1":      []*Box{{2, 2, 6, 6}},
		"WEAPON_SHIP_NORMAL_1": []*Box{{2, 2, 6, 6}},
		"laser1_0":             []*Box{{5, 1, 10, 2}},
		"laser1_15":            []*Box{{3, 2, 5, 1}, {6, 3, 7, 1}, {7, 4, 8, 1}},
		"laser1_30":            []*Box{{12, 7, 3, 3}, {7, 4, 3, 3}},
		"laser1_45":            []*Box{{7, 6, 3, 3}, {10, 10, 3, 2}},
		"laser1_60":            []*Box{{5, 7, 3, 3}, {8, 12, 3, 3}},
		"laser1_75":            []*Box{{4, 8, 2, 2}, {4, 11, 3, 4}},
		"laser1_90":            []*Box{{13, 4, 2, 10}},
		"laser1_105":           []*Box{{10, 12, 2, 3}, {11, 6, 2, 6}},
		"laser1_120":           []*Box{{6, 10, 2, 5}, {9, 6, 2, 5}},
		"laser1_135":           []*Box{{3, 10, 3, 3}, {7, 7, 3, 3}},
		"laser1_150":           []*Box{{1, 7, 3, 3}, {5, 5, 3, 3}},
		"laser1_165":           []*Box{{1, 4, 3, 3}, {6, 3, 3, 3}},
		"laser1_180":           []*Box{{1, 13, 13, 2}},
		"laser1_195":           []*Box{{1, 10, 3, 2}, {4, 11, 6, 2}},
		"laser1_210":           []*Box{{1, 6, 3, 3}, {6, 8, 3, 3}},
		"laser1_225":           []*Box{{3, 3, 3, 3}, {7, 7, 3, 3}},
		"laser1_240":           []*Box{{5, 1, 3, 3}, {7, 5, 3, 3}},
		"laser1_255":           []*Box{{9, 1, 3, 3}, {10, 6, 3, 3}},
		"laser1_270":           []*Box{{1, 1, 2, 9}},
		"laser1_285":           []*Box{{4, 1, 2, 3}, {3, 5, 2, 6}},
		"laser1_300":           []*Box{{7, 1, 3, 3}, {4, 6, 3, 3}},
		"laser1_315":           []*Box{{9, 3, 3, 3}, {5, 7, 3, 3}},
		"laser1_330":           []*Box{{12, 6, 3, 3}, {8, 8, 3, 3}},
		"laser1_345":           []*Box{{7, 10, 3, 3}, {12, 9, 3, 3}},
		"E_SENKAN_1_3":         []*Box{{5, 15, 64, 150}, {17, 165, 40, 30}},
		"E_ROBO1":              []*Box{{0, 0, 24, 24}},
		"E_SENKAN_1_1":         []*Box{{36, 13, 42, 52}, {16, 65, 89, 68}},
		"E_SENKAN_1_2":         []*Box{{14, 16, 159, 44}},
		"E_SENKAN_1_4":         []*Box{{15, 20, 167, 36}},
		"E_ROBO2":              []*Box{{0, 0, 48, 48}},
		"E_COLONY1":            []*Box{{85, 21, 88, 131}, {23, 159, 216, 107}, {63, 268, 129, 81}},
		"SHOT_BIG1":            []*Box{{3, 3, 13, 13}},
		"P_ROBO_1":             []*Box{{7, 7, 2, 2}},
		"P_ROBO_2":             []*Box{{9, 9, 6, 6}},
		"P_ROBO_3":             []*Box{{7, 7, 2, 2}},
		"P_ROBO_4":             []*Box{{12, 12, 4, 4}},
		"P_SHIP1":              []*Box{{28, 23, 105, 32}, {28, 22, 70, 7}},
		"E_ROBO4":              []*Box{{4, 4, 12, 12}},
		"P_ROBO_8":             []*Box{{11, 11, 10, 10}},
		"E_BOSS":               []*Box{{21, 50, 157, 28}, {137, 77, 30, 24}, {166, 82, 29, 11}, {42, 34, 164, 16}, {64, 21, 99, 13}},
		"PE_ROBO1":             []*Box{{7, 7, 9, 9}},
		"PE_ROBO2":             []*Box{{16, 16, 16, 16}},
		"PE_SHOOTER":           []*Box{{15, 15, 19, 19}},
		"PE_SENKAN_1_4":        []*Box{{54, 20, 117, 8}, {17, 30, 162, 10}, {32, 41, 148, 12}},
		"BARRIAR":              []*Box{{6, 6, 48, 48}},
		"E_BOSS2":              []*Box{{28, 72, 122, 77}, {74, 26, 38, 139}},
		"E_ROBO9":              []*Box{{12, 12, 4, 4}},
		"E_BOSS3":              []*Box{{37, 33, 100, 95}, {77, 7, 25, 23}},
		"E_BOSS4":              []*Box{{25, 82, 80, 44}, {48, 130, 30, 18}, {62, 35, 6, 44}},
		"SMALL_BULLET":         []*Box{{2, 2, 3, 3}},
		"KAITEN_BULLET":        []*Box{{2, 2, 4, 4}},
		"E_ZAKO":               []*Box{{4, 4, 12, 12}},
		"ITEM_LIFE":            []*Box{{0, 0, 16, 10}},
		"ITEM_P":               []*Box{{0, 0, 16, 17}},
		"NULL":                 []*Box{{0, 0, 0, 0}},
		"NUCLEAR":              []*Box{{6, 5, 36, 38}, {12, 2, 24, 44}, {1, 14, 46, 20}},
	}
)

// GetCollisionBox returns collision box
func GetCollisionBox(kind string) []*Box {
	box := boxData[kind]
	if box == nil {
		panic(fmt.Sprintf("invalid collision box kind: %s", kind))
	}
	return box
}

// IsCollideWith returns if it collides with another actor
func IsCollideWith(c1, c2 Collider) bool {
	x1 := c1.GetX() - c1.GetWidth()/2
	y1 := c1.GetY() - c1.GetHeight()/2
	x2 := c2.GetX() - c2.GetWidth()/2
	y2 := c2.GetY() - c2.GetHeight()/2

	list1 := c1.GetCollisionBox()
	list2 := c2.GetCollisionBox()

	for _, b1 := range list1 {
		for _, b2 := range list2 {
			if isBoxCollide(x1, y1, x2, y2, b1, b2) {
				return true
			}
		}
	}
	return false
}

func isBoxCollide(x1, y1, x2, y2 float64, b1, b2 *Box) bool {
	return x1+b1.x <= x2+b2.x+b2.w &&
		x2+b2.x <= x1+b1.x+b1.w &&
		y1+b1.y <= y2+b2.y+b2.h &&
		y2+b2.y <= y1+b1.y+b1.h
}
