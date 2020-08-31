package movable

// Boundarizer represents the boundarizer of the field the ShObject can move inside
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

// ShObject represents the base of player, enemy, shots
type ShObject struct {
	x      float64
	y      float64
	speed  float64
	vx     float64
	vy     float64
	width  int
	height int
	deg    int
}

// SetSpeed sets the speed
func (a *ShObject) SetSpeed(speed float64) {
	a.speed = speed
}

// SetPostion sets the position
func (a *ShObject) SetPosition(x, y float64) {
	a.x = x
	a.y = y
}

// SetSize sets the size
func (a *ShObject) SetSize(width, height int) {
	a.width = width
	a.height = height
}

// GetPosition returns the position in (x, y)
func (a *ShObject) GetPosition() (float64, float64) {
	return a.x, a.y
}

// GetDeg returns the degree of the ShObject
func (a *ShObject) GetDeg() int {
	return a.deg
}

// GetNormalizedDegree returns normalized degree
func (a *ShObject) GetNormalizedDegree() int {
	adjust := 22.5
	return int((float64(((a.deg+360)%360))+adjust)/45) * 45
}

// GetX returns x
func (a *ShObject) GetX() int {
	return int(a.x)
}

// GetY returns y
func (a *ShObject) GetY() int {
	return int(a.y)
}

// GetWidth returns width
func (a *ShObject) GetWidth() int {
	return a.width
}

// GetHeight returns height
func (a *ShObject) GetHeight() int {
	return a.height
}

func (a *ShObject) isOutOfBoundary() bool {
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
