package actor

import (
	"math"
)

var (
	directionDegreeMap = map[int]int{1: 135, 2: 90, 3: 45, 4: 180, 6: 0, 7: 225, 8: 270, 9: 315}
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
	a.NSpd = math.Cos(45*math.Pi/180) * a.Speed
}

// SetPosition sets the position
func (a *Actor) SetPosition(x, y float64) {
	a.X = x
	a.Y = y
}

// SetDirection sets the direction
func (a *Actor) SetDirection(direction int) {
	a.Direction = direction
	a.DirectionDegree = directionDegreeMap[direction]
}
