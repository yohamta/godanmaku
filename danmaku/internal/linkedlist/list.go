package linkedlist

import (
	"unsafe"
)

type List struct {
	length int
	first  *Element
	last   *Element
	ite    *Iterator
}

func NewList() *List {
	l := &List{}
	l.ite = &Iterator{}

	return l
}

func (l *List) Length() int {
	return l.length
}

func (l *List) AddValue(value unsafe.Pointer) {
	e := NewElement(value)
	l.AddElement(e)
}

func (l *List) GetFirstElement() *Element {
	return l.first
}

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

func (l *List) Clear() {
	l.first = nil
	l.last = nil
	l.length = 0
}

func (l *List) GetIterator() *Iterator {
	l.ite.list = l
	l.ite.current = l.GetFirstElement()
	return l.ite
}
