package quad

import (
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/list"
)

type Object interface {
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
}

type Quad struct {
	x1, x2, y1, y2 float64
	ne             *Quad
	nw             *Quad
	se             *Quad
	sw             *Quad
	isLeaf         bool
	nodes          *list.List
	ite            *Iterator
	descendants    []*Quad
}

type Node struct {
	x1, x2, y1, y2 float64
	ptr            unsafe.Pointer
	element        *list.Element
	quad           *Quad
}

func NewNode(p unsafe.Pointer) *Node {
	n := &Node{}
	n.ptr = p
	n.element = list.NewElement(unsafe.Pointer(n))
	return n
}

func (n *Node) GetItem() unsafe.Pointer {
	return n.ptr
}

func (n *Node) SetItem(p unsafe.Pointer) {
	n.ptr = p
}

func NewQuad(x1, x2, y1, y2 float64, depth int) *Quad {
	q := &Quad{}

	q.x1 = x1
	q.x2 = x2
	q.y1 = y1
	q.y2 = y2
	q.nodes = list.NewList()
	q.descendants = []*Quad{q}
	q.ite = &Iterator{quad: q}

	if depth == 1 {
		q.isLeaf = true
	} else {
		q.nw = NewQuad(x1, x1+(x2-x1)/2, y1, y1+(y2-y1)/2, depth-1)
		q.ne = NewQuad(x1+(x2-x1)/2, x2, y1, y1+(y2-y1)/2, depth-1)
		q.sw = NewQuad(x1, x1+(x2-x1)/2, y1+(y2-y1)/2, y2, depth-1)
		q.se = NewQuad(x1+(x2-x1)/2, x2, y1+(y2-y1)/2, y2, depth-1)
		q.descendants = append(q.descendants, q.nw.descendants...)
		q.descendants = append(q.descendants, q.ne.descendants...)
		q.descendants = append(q.descendants, q.sw.descendants...)
		q.descendants = append(q.descendants, q.se.descendants...)
	}

	return q
}

func (q *Quad) GetIterator() *Iterator {
	q.ite.index = 0
	q.ite.current = q.nodes.GetIterator()

	return q.ite
}

func (q *Quad) SearchQuad(o Object) *Quad {
	x1 := o.GetX() - o.GetWidth()/2
	x2 := o.GetX() + o.GetWidth()/2
	y1 := o.GetY() - o.GetHeight()/2
	y2 := o.GetY() + o.GetHeight()/2
	return findQuad(q, x1, x2, y1, y2)
}

func isContain(q *Quad, x1, x2, y1, y2 float64) bool {
	return q.x1 <= x1 && x2 <= q.x2 && q.y1 <= y1 && y2 <= q.y2
}

func findQuad(root *Quad, x1, x2, y1, y2 float64) *Quad {
	var q *Quad = root
	for true {
		if q.isLeaf {
			break
		}
		if x2 <= q.x1+(q.x2-q.x1)/2 && y2 <= q.y1+(q.y2-q.y1)/2 {
			if isContain(q.nw, x1, x2, y1, y2) {
				q = q.nw
				continue
			}
		} else if x1 >= q.x2 && y2 <= q.y1+(q.y2-q.y1)/2 {
			if isContain(q.ne, x1, x2, y1, y2) {
				q = q.ne
				continue
			}
		} else if x2 <= q.x1+(q.x2-q.x1)/2 {
			if isContain(q.sw, x1, x2, y1, y2) {
				q = q.sw
				continue
			}
		} else {
			if isContain(q.se, x1, x2, y1, y2) {
				q = q.se
				continue
			}
		}
		break
	}
	return q
}

func RemoveNodeFromQuad(node *Node) {
	if node.quad == nil {
		return
	}
	node.quad.nodes.RemoveElement(node.element)
	node.quad = nil
}

func (q *Quad) AddNode(o Object, node *Node) {
	RemoveNodeFromQuad(node)

	x1 := o.GetX() - o.GetWidth()/2
	x2 := o.GetX() + o.GetWidth()/2
	y1 := o.GetY() - o.GetHeight()/2
	y2 := o.GetY() + o.GetHeight()/2

	node.x1 = x1
	node.x2 = x2
	node.y1 = y1
	node.y2 = y2

	found := findQuad(q, x1, x2, y1, y2)

	node.quad = found
	found.nodes.AddElement(node.element)
}
