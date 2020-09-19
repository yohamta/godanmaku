package quad

import (
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/list"
)

// Object represents obejct to be contained
type Object interface {
	getX() float64
	getY() float64
	getWidth() float64
	getHeight() float64
}

// Quad represents quad tree
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

// Node represents a node that a quad contains
type Node struct {
	x1, x2, y1, y2 float64
	object         Object
	ptr            unsafe.Pointer
	element        *list.Element
	quad           *Quad
}

// NewNode creates new ndoe
func NewNode(o Object, p unsafe.Pointer) *Node {
	n := &Node{}
	n.object = o
	n.ptr = p
	n.element = list.NewElement(unsafe.Pointer(n))
	return n
}

// GetItem returns the pointer of the item
func (n *Node) GetItem() unsafe.Pointer {
	return n.ptr
}

// NewQuad creates new quad
func NewQuad(x1, x2, y1, y2 float64, depth int) *Quad {
	q := &Quad{}

	q.x1 = x1
	q.x2 = x2
	q.y1 = y1
	q.y2 = y2
	q.nodes = list.NewList()
	q.descendants = []*Quad{q}
	q.ite = &Iterator{quad: q}

	if depth == 0 {
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

// GetIterator returns iterator
func (q *Quad) GetIterator() *Iterator {
	q.ite.index = 0
	q.ite.current = q.nodes.GetIterator()

	return q.ite
}

// Search returns quad the object belongs to
func (q *Quad) Search(o Object) *Quad {
	x1 := o.getX() - o.getWidth()/2
	x2 := o.getX() + o.getWidth()/2
	y1 := o.getY() - o.getHeight()/2
	y2 := o.getY() + o.getHeight()/2
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

func removeFromQuad(node *Node) {
	if node.quad == nil {
		return
	}
	node.quad.nodes.RemoveElement(node.element)
	node.quad = nil
}

// RemoveNode removes a node from the quad
func (q *Quad) RemoveNode(node *Node) {
	removeFromQuad(node)
}

// AddNode adds a node to the quad
func (q *Quad) AddNode(node *Node) {
	removeFromQuad(node)

	o := node.object
	x1 := o.getX() - o.getWidth()/2
	x2 := o.getX() + o.getWidth()/2
	y1 := o.getY() - o.getHeight()/2
	y2 := o.getY() + o.getHeight()/2
	node.x1 = x1
	node.x2 = x2
	node.y1 = y1
	node.y2 = y2

	found := findQuad(q, x1, x2, y1, y2)
	node.quad = found
	found.nodes.AddElement(node.element)
}
