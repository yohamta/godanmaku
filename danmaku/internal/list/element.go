package list

import (
	"unsafe"
)

type Value interface{}

type Element struct {
	value unsafe.Pointer
	prev  *Element
	next  *Element
}

func NewElement(value unsafe.Pointer) *Element {
	e := &Element{}
	e.value = value

	return e
}

func (e *Element) GetValue() unsafe.Pointer {
	return e.value
}

func (e *Element) GetNext() *Element {
	return e.next
}
