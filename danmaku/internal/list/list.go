package list

// List represents container
type List struct {
	length int
	first  *Element
	last   *Element
}

// NewList creates new element
func NewList() *List {
	l := &List{}

	return l
}

// AddValue returns the value
func (l *List) AddValue(v Value) {
	e := NewElement(v)
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
