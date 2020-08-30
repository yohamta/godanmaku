package actors

import "math"

// Actor represents actor
type Actor struct {
	X      float64
	Y      float64
	Speed  float64
	Vx     float64
	Vy     float64
	Width  int
	Height int
	Deg    int
}

// SetSpeed sets the speed to the actor
func (a *Actor) SetSpeed(speed float64) {
	a.Speed = speed
}

// SetPosition sets the position
func (a *Actor) SetPosition(x, y float64) {
	a.X = x
	a.Y = y
}

// SetDeg sets the direction in degree
func (a *Actor) SetDeg(degree int) {
	a.Deg = degree
}

// DegreeToDirectionIndex converts degree to direction index
func DegreeToDirectionIndex(degree int) int {
	adjust := 22.5
	return int(float64(degree)+90.0+360.0+adjust) % 360 / 45
}

// RadToDeg converts degree to radian
func RadToDeg(radian float64) int {
	return int(radian * 180 / math.Pi)
}

// DegToRad converts degree to radian
func DegToRad(degree int) float64 {
	return float64(degree) * math.Pi / 180
}
