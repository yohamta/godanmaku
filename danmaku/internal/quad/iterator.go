package quad

import (
	"github.com/yohamta/godanmaku/danmaku/internal/list"
)

// Iterator represents iterator
type Iterator struct {
	quad    *Quad
	current *list.Iterator
	index   int
}

// HasNext returns if this has next element
func (ite *Iterator) HasNext() bool {
	if ite.current.HasNext() {
		return true
	}
	if ite.index >= len(ite.quad.descendants) {
		return false
	}
	for ite.index < len(ite.quad.descendants) {
		ite.index++
		if ite.quad.descendants[ite.index].nodes.Length() > 0 {
			ite.current = ite.quad.descendants[ite.index].nodes.GetIterator()
			return true
		}
	}
	return false
}

// Next returns next element
func (ite *Iterator) Next() *Node {
	if ite.HasNext() == false {
		panic("something went wrong in quad iterator")
	}
	n := (*Node)(ite.current.Next().GetValue())
	return n
}
