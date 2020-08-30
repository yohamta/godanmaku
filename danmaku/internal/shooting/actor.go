package shooting

import "math"

// Boundary represents the boundary of the field the actor can move inside
type Boundary interface {
	GetLeft() int
	GetTop() int
	GetRight() int
	GetBottom() int
}

// Actor represents the base of player, enemy, shots
type Actor struct {
	x      float64
	y      float64
	speed  float64
	vx     float64
	vy     float64
	width  int
	height int
	deg    int
}

func (a *Actor) setSpeed(speed float64) {
	a.speed = speed
}

func (a *Actor) setPosition(x, y float64) {
	a.x = x
	a.y = y
}

func (a *Actor) setDeg(degree int) {
	a.deg = degree
}

func (a *Actor) getCenter() (int, int) {
	x := int(a.x)
	y := int(a.y)
	return x, y
}

func (a *Actor) isOutOfBoundary(boundary Boundary) bool {
	if int(a.x)+a.width/2 < boundary.GetLeft() {
		return true
	}
	if int(a.x)-a.width > boundary.GetRight() {
		return true
	}
	if int(a.y)+a.height < boundary.GetTop() {
		return true
	}
	if int(a.y)-a.height/2 > boundary.GetBottom() {
		return true
	}
	return false
}

func degreeToDirectionIndex(degree int) int {
	adjust := 22.5
	return int(float64(degree)+90.0+360.0+adjust) % 360 / 45
}

func radToDeg(radian float64) int {
	return int(radian * 180 / math.Pi)
}

func degToRad(degree int) float64 {
	return float64(degree) * math.Pi / 180
}
