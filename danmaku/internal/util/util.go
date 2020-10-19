package util

import (
	"image"
	"math"

	"github.com/yohamta/godanmaku/danmaku/internal/shared"
)

// DegreeToDirectionIndex convert degree into 1 to 8 integer
func DegreeToDirectionIndex(degree int) int {
	adjust := 22.5
	return int(float64(degree)+90.0+360.0+adjust) % 360 / 45
}

// RadToDeg converts radian to degree
func RadToDeg(radian float64) int {
	return int(radian * 180 / math.Pi)
}

// DegToRad converts degree to radian
func DegToRad(degree int) float64 {
	return float64(degree) * math.Pi / 180
}

// Entity represents an entity
type Entity interface {
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
}

// Area represents an area
type Area interface {
	GetLeft() float64
	GetRight() float64
	GetTop() float64
	GetBottom() float64
}

// IsOutOfArea Returns if the entity is out of the certain area
func IsOutOfArea(e Entity, area Area) bool {
	if e.GetX()+e.GetWidth()/2 < area.GetLeft() {
		return true
	}
	if e.GetX()-e.GetWidth() > area.GetRight() {
		return true
	}
	if e.GetY()+e.GetHeight() < area.GetTop() {
		return true
	}
	if e.GetY()-e.GetHeight()/2 > area.GetBottom() {
		return true
	}
	return false
}

// IsOutOfAreaEnoughly Returns if the entity is enoughly out of the certain area
func IsOutOfAreaEnoughly(e Entity, area Area) bool {
	w := float64(shared.ScreenSize.X)
	h := float64(shared.ScreenSize.Y)
	if e.GetX()+e.GetWidth()/2 < area.GetLeft()-w/2 {
		return true
	}
	if e.GetX()-e.GetWidth() > area.GetRight()+w/2 {
		return true
	}
	if e.GetY()+e.GetHeight() < area.GetTop()-h/2 {
		return true
	}
	if e.GetY()-e.GetHeight()/2 > area.GetBottom()+h/2 {
		return true
	}
	return false
}

func PrintRect(rect image.Rectangle) {
	println(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
}
