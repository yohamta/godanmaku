package collision

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
