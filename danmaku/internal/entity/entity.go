package entity

import "github.com/yohamta/godanmaku/danmaku/internal/field"

// Area represents certain area
type Area interface {
	getPosition() (float64, float64)
	getSize() (float64, float64)
}

// Entity represents an entity
type Entity struct {
	x, y          float64
	width, height float64
	currField     *field.Field
	isActive      bool
}

// NewEntity creates new Entity
func NewEntity() *Entity {
	e := &Entity{}
	e.isActive = false
	return e
}

// IsActive returns if the actor is active in bool
func (e *Entity) IsActive() bool {
	return e.isActive
}

// SetActive returns if the actor is active in bool
func (e *Entity) SetActive(isActive bool) {
	e.isActive = isActive
}

// SetSize returns the size
func (e *Entity) SetSize(width, height float64) {
	e.width = width
	e.height = height
}

// SetPosition sets the position
func (e *Entity) SetPosition(x, y float64) {
	e.x = x
	e.y = y
}

// GetPosition returns the position
func (e *Entity) GetPosition() (float64, float64) {
	return e.x, e.y
}

// GetX returns x
func (e *Entity) GetX() float64 {
	return e.x
}

// GetY returns y
func (e *Entity) GetY() float64 {
	return e.y
}

// GetSize returns the size
func (e *Entity) GetSize() (float64, float64) {
	return e.width, e.height
}

// GetWidth returns width
func (e *Entity) GetWidth() float64 {
	return e.width
}

// GetHeight returns height
func (e *Entity) GetHeight() float64 {
	return e.height
}

// GetField returns field
func (e *Entity) GetField() *field.Field {
	return e.currField
}

// SetField returns field
func (e *Entity) SetField(f *field.Field) {
	e.currField = f
}

// IsOutOfField Returns if the entity is out of the certain area
func (e *Entity) IsOutOfField() bool {
	f := e.currField
	if e.x+e.width/2 < f.GetLeft() {
		return true
	}
	if e.x-e.width > f.GetRight() {
		return true
	}
	if e.y+e.height < f.GetTop() {
		return true
	}
	if e.y-e.height/2 > f.GetBottom() {
		return true
	}
	return false
}

// IsCollideWith returns if it collides with another actor
func (e *Entity) IsCollideWith(other *Entity) bool {
	return e.GetX() <= other.GetX()+other.GetWidth() &&
		other.GetX() <= e.GetX()+e.GetWidth() &&
		e.GetY() <= other.GetY()+other.GetHeight() &&
		other.GetY() <= e.GetY()+e.GetHeight()
}
