package list

// Iterator represents iterator
type Iterator struct {
	list    *List
	current *Element
}

// HasNext returns if this has next element
func (ite *Iterator) HasNext() bool {
	return ite.current != nil
}

// Next returns next element
func (ite *Iterator) Next() *Element {
	e := ite.current
	ite.current = e.GetNext()
	return e
}
