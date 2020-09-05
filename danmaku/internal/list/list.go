package list

import (
	"unsafe"
)

// List represents container
type List struct {
	length int
	first  *Element
	last   *Element
	ite    *Iterator
}

// NewList creates new element
func NewList() *List {
	l := &List{}
	l.ite = &Iterator{}

	return l
}

// AddValue returns the value
func (l *List) AddValue(value unsafe.Pointer) {
	e := NewElement(value)
	l.AddElement(e)
}

// GetFirstElement returns the fist element
func (l *List) GetFirstElement() *Element {
	return l.first
}

// RemoveElement removes the element
func (l *List) RemoveElement(e *Element) {
	prev := e.prev
	next := e.next
	if prev == nil {
		l.first = next
	} else {
		prev.next = next
	}
	if next == nil {
		l.last = prev
	} else {
		next.prev = prev
	}
	l.length--
}

// AddElement adds new element
func (l *List) AddElement(e *Element) {
	e.prev = nil
	e.next = nil
	if l.length == 0 {
		l.first = e
		l.last = e
	} else {
		l.last.next = e
		e.prev = l.last
		l.last = e
	}
	l.length++
}

// Clear adds new element
func (l *List) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

// GetIterator retusn iterator
func (l *List) GetIterator() *Iterator {
	l.ite.list = l
	l.ite.current = l.GetFirstElement()
	return l.ite
}
