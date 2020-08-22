package actor

import (
	"math"
)

// Actor represents actor
type Actor struct {
	X               float64
	Y               float64
	Speed           float64
	NSpd            float64
	Width           int
	Height          int
	Direction       int
	DirectionDegree int
}

// SetSpeed sets the speed to the actor
func (a *Actor) SetSpeed(speed float64) {
	a.Speed = speed
	a.NSpd = math.Cos(45 * math.Pi / 180)
}

// SetPosition sets the position
func (a *Actor) SetPosition(x, y float64) {
	a.X = x
	a.Y = y
}
