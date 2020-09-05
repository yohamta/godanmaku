package flyweight

import (
	"github.com/yohamta/godanmaku/danmaku/internal/list"
)

// Iterator represents iterator
type Iterator struct {
	current *list.Element
}

// HasNext returns if this has next element
func (ite *Iterator) HasNext() bool {
	return ite.current != nil
}

// Next returns next element
func (ite *Iterator) Next() *Object {
	e := ite.current
	ite.current = e.GetNext()

	return (*Object)(e.GetValue())
}
