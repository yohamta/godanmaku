package quad

import (
	"github.com/yohamta/godanmaku/danmaku/internal/linkedlist"
)

type Iterator struct {
	quad    *Quad
	current *linkedlist.Iterator
	index   int
}

func (ite *Iterator) HasNext() bool {
	if ite.current.HasNext() {
		return true
	}
	if ite.index >= len(ite.quad.descendants) {
		return false
	}
	for ite.index < len(ite.quad.descendants)-1 {
		ite.index++
		if ite.quad.descendants[ite.index].nodes.Length() > 0 {
			ite.current = ite.quad.descendants[ite.index].nodes.GetIterator()
			return true
		}
	}
	return false
}

func (ite *Iterator) Next() *Node {
	if ite.HasNext() == false {
		panic("something went wrong in quad iterator")
	}
	n := (*Node)(ite.current.Next().GetValue())
	return n
}
