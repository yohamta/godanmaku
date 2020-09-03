package list

// Value represents value to be stored
type Value interface{}

// Element represents container
type Element struct {
	value Value
	prev  *Element
	next  *Element
}

// NewElement creates new element
func NewElement(v Value) *Element {
	e := &Element{value: v}

	return e
}

// GetValue returns the value
func (e *Element) GetValue() Value {
	return e.value
}

// GetNext returns the next value
func (e *Element) GetNext() *Element {
	return e.next
}
