package list

import (
	"unsafe"
)

// Value represents value to be stored
type Value interface{}

// Element represents container
type Element struct {
	value unsafe.Pointer
	prev  *Element
	next  *Element
}

// NewElement creates new element
func NewElement(value unsafe.Pointer) *Element {
	e := &Element{}
	e.value = value

	return e
}

// GetValue returns the value
func (e *Element) GetValue() unsafe.Pointer {
	return e.value
}

// GetNext returns the next value
func (e *Element) GetNext() *Element {
	return e.next
}
