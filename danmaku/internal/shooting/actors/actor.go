package actors

import "math"

// Boundarizer represents the boundarizer of the field the actor can move inside
type Boundarizer interface {
	GetLeft() int
	GetTop() int
	GetRight() int
	GetBottom() int
}

var (
	boundarizer Boundarizer
)

// SetBoundary sets the boundarizer of the field
func SetBoundary(b Boundarizer) {
	boundarizer = b
}

type position struct {
	x, y float64
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

// GetPosition returns the position in (x, y)
func (a *Actor) GetPosition() (float64, float64) {
	return a.x, a.y
}

// GetDeg returns the degree of the actor
func (a *Actor) GetDeg() int {
	return a.deg
}

// GetNormalizedDegree returns normalized degree
func (a *Actor) GetNormalizedDegree() int {
	adjust := 22.5
	return int((float64(((a.deg+360)%360))+adjust)/45) * 45
}

func (a *Actor) isOutOfBoundary() bool {
	if int(a.x)+a.width/2 < boundarizer.GetLeft() {
		return true
	}
	if int(a.x)-a.width > boundarizer.GetRight() {
		return true
	}
	if int(a.y)+a.height < boundarizer.GetTop() {
		return true
	}
	if int(a.y)-a.height/2 > boundarizer.GetBottom() {
		return true
	}
	return false
}

// GetX returns x
func (a *Actor) GetX() int {
	return int(a.x)
}

// GetY returns y
func (a *Actor) GetY() int {
	return int(a.y)
}

// GetWidth returns width
func (a *Actor) GetWidth() int {
	return a.width
}

// GetHeight returns height
func (a *Actor) GetHeight() int {
	return a.height
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

// Collider represents collidable struct
type Collider interface {
	GetX() int
	GetY() int
	GetWidth() int
	GetHeight() int
}

// IsCollideWith returns if it collides with another actor
func IsCollideWith(c1 Collider, c2 Collider) bool {
	return c1.GetX() <= c2.GetX()+c2.GetWidth() &&
		c2.GetX() <= c1.GetX()+c1.GetWidth() &&
		c1.GetY() <= c2.GetY()+c2.GetHeight() &&
		c2.GetY() <= c1.GetY()+c1.GetHeight()
}
