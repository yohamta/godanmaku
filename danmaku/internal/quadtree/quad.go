package quadtree

import (
	"unsafe"

	"github.com/miyahoyo/godanmaku/danmaku/internal/linkedlist"
)

type Collider interface {
	GetRect() (x0 float64, y0 float64, x1 float64, y1 float64)
}

type Quadtree struct {
	x0, y0, x1, y1 float64
	ne             *Quadtree
	nw             *Quadtree
	se             *Quadtree
	sw             *Quadtree
	isLeaf         bool
	nodes          *linkedlist.List
	ite            *Iterator
	descendants    []*Quadtree
}

type Node struct {
	x0, y0, x1, y1 float64
	ptr            unsafe.Pointer
	element        *linkedlist.Element
	quad           *Quadtree
}

func NewNode(p unsafe.Pointer) *Node {
	n := &Node{}
	n.ptr = p
	n.element = linkedlist.NewElement(unsafe.Pointer(n))
	return n
}

func (n *Node) GetItem() unsafe.Pointer {
	return n.ptr
}

func (n *Node) SetItem(p unsafe.Pointer) {
	n.ptr = p
}

func NewQuadtree(x0, y0, x1, y1 float64, depth int) *Quadtree {
	q := &Quadtree{}

	q.x0 = x0
	q.y0 = y0
	q.x1 = x1
	q.y1 = y1
	q.nodes = linkedlist.NewList()
	q.descendants = []*Quadtree{q}
	q.ite = &Iterator{quad: q}

	if depth == 1 {
		q.isLeaf = true
	} else {
		q.nw = NewQuadtree(x0, y0, x0+(x1-x0)/2, y0+(y1-y0)/2, depth-1)
		q.ne = NewQuadtree(x0+(x1-x0)/2, y0, x1, y0+(y1-y0)/2, depth-1)
		q.sw = NewQuadtree(x0, y0+(y1-y0)/2, x0+(x1-x0)/2, y1, depth-1)
		q.se = NewQuadtree(x0+(x1-x0)/2, y0+(y1-y0)/2, x1, y1, depth-1)
		q.descendants = append(q.descendants, q.nw.descendants...)
		q.descendants = append(q.descendants, q.ne.descendants...)
		q.descendants = append(q.descendants, q.sw.descendants...)
		q.descendants = append(q.descendants, q.se.descendants...)
	}

	return q
}

func (q *Quadtree) GetIterator() *Iterator {
	q.ite.index = 0
	q.ite.current = q.nodes.GetIterator()

	return q.ite
}

func (q *Quadtree) SearchQuadtree(c Collider) *Quadtree {
	x0, y0, x1, y1 := c.GetRect()
	return findQuadtree(q, x0, x1, y0, y1)
}

func isContain(q *Quadtree, x0, x1, y0, y1 float64) bool {
	return q.x0 <= x0 && x1 <= q.x1 && q.y0 <= y0 && y1 <= q.y1
}

func findQuadtree(root *Quadtree, x0, x1, y0, y1 float64) *Quadtree {
	var q *Quadtree = root
	for true {
		if q.isLeaf {
			break
		}
		if x1 <= q.x0+(q.x1-q.x0)/2 && y1 <= q.y0+(q.y1-q.y0)/2 {
			if isContain(q.nw, x0, x1, y0, y1) {
				q = q.nw
				continue
			}
		} else if x0 >= q.x1 && y1 <= q.y0+(q.y1-q.y0)/2 {
			if isContain(q.ne, x0, x1, y0, y1) {
				q = q.ne
				continue
			}
		} else if x1 <= q.x0+(q.x1-q.x0)/2 {
			if isContain(q.sw, x0, x1, y0, y1) {
				q = q.sw
				continue
			}
		} else {
			if isContain(q.se, x0, x1, y0, y1) {
				q = q.se
				continue
			}
		}
		break
	}
	return q
}

func RemoveNodeFromQuadtree(node *Node) {
	if node.quad == nil {
		return
	}
	node.quad.nodes.RemoveElement(node.element)
	node.quad = nil
}

func (q *Quadtree) AddNode(c Collider, node *Node) {
	RemoveNodeFromQuadtree(node)
	x0, y0, x1, y1 := c.GetRect()

	node.x0 = x0
	node.x1 = x1
	node.y0 = y0
	node.y1 = y1

	found := findQuadtree(q, x0, x1, y0, y1)

	node.quad = found
	found.nodes.AddElement(node.element)
}
