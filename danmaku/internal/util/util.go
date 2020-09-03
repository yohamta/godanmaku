package util

import "math"

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

// IsCollideWith returns if it collides with another actor
func IsCollideWith(e1 Entity, e2 Entity) bool {
	return e1.GetX() <= e2.GetX()+e2.GetWidth() &&
		e2.GetX() <= e1.GetX()+e1.GetWidth() &&
		e1.GetY() <= e2.GetY()+e2.GetHeight() &&
		e2.GetY() <= e1.GetY()+e1.GetHeight()
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
