package quadtree

import (
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/linkedlist"
)

type Collidar interface {
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
}

type Quadtree struct {
	x1, x2, y1, y2 float64
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
	x1, x2, y1, y2 float64
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

func NewQuadtree(x1, x2, y1, y2 float64, depth int) *Quadtree {
	q := &Quadtree{}

	q.x1 = x1
	q.x2 = x2
	q.y1 = y1
	q.y2 = y2
	q.nodes = linkedlist.NewList()
	q.descendants = []*Quadtree{q}
	q.ite = &Iterator{quad: q}

	if depth == 1 {
		q.isLeaf = true
	} else {
		q.nw = NewQuadtree(x1, x1+(x2-x1)/2, y1, y1+(y2-y1)/2, depth-1)
		q.ne = NewQuadtree(x1+(x2-x1)/2, x2, y1, y1+(y2-y1)/2, depth-1)
		q.sw = NewQuadtree(x1, x1+(x2-x1)/2, y1+(y2-y1)/2, y2, depth-1)
		q.se = NewQuadtree(x1+(x2-x1)/2, x2, y1+(y2-y1)/2, y2, depth-1)
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

func (q *Quadtree) SearchQuadtree(c Collidar) *Quadtree {
	x1 := c.GetX() - c.GetWidth()/2
	x2 := c.GetX() + c.GetWidth()/2
	y1 := c.GetY() - c.GetHeight()/2
	y2 := c.GetY() + c.GetHeight()/2
	return findQuadtree(q, x1, x2, y1, y2)
}

func isContain(q *Quadtree, x1, x2, y1, y2 float64) bool {
	return q.x1 <= x1 && x2 <= q.x2 && q.y1 <= y1 && y2 <= q.y2
}

func findQuadtree(root *Quadtree, x1, x2, y1, y2 float64) *Quadtree {
	var q *Quadtree = root
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

func RemoveNodeFromQuadtree(node *Node) {
	if node.quad == nil {
		return
	}
	node.quad.nodes.RemoveElement(node.element)
	node.quad = nil
}

func (q *Quadtree) AddNode(c Collidar, node *Node) {
	RemoveNodeFromQuadtree(node)

	x1 := c.GetX() - c.GetWidth()/2
	x2 := c.GetX() + c.GetWidth()/2
	y1 := c.GetY() - c.GetHeight()/2
	y2 := c.GetY() + c.GetHeight()/2

	node.x1 = x1
	node.x2 = x2
	node.y1 = y1
	node.y2 = y2

	found := findQuadtree(q, x1, x2, y1, y2)

	node.quad = found
	found.nodes.AddElement(node.element)
}
