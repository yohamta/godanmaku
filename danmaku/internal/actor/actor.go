package actor

// Actor represents actor
type Actor struct {
	X      float64
	Y      float64
	Speed  float64
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
