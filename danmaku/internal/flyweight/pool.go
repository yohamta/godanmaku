package flyweight

import (
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/list"
)

// Item represents item
type Item interface {
	IsActive() bool
}

// Pool represents factory
type Pool struct {
	actives *list.List
	pool    *list.List
	ite     *Iterator
}

// NewPool creates new Pool
func NewPool() *Pool {
	p := &Pool{}
	p.actives = list.NewList()
	p.pool = list.NewList()
	p.ite = &Iterator{}

	return p
}

// GetActiveNum returns the number of active item
func (p *Pool) GetActiveNum() int {
	return p.actives.Length()
}

// AddToPool adds resusable item
func (p *Pool) AddToPool(item unsafe.Pointer) {
	o := &Object{}
	o.data = item
	o.isActive = false
	ptr := unsafe.Pointer(o)
	elem := list.NewElement(ptr)
	o.elem = elem
	p.pool.AddElement(elem)
}

// GetIterator adds resusable item
func (p *Pool) GetIterator() *Iterator {
	ite := p.ite
	ite.current = p.actives.GetFirstElement()

	return ite
}

// CreateFromPool adds resusable item
func (p *Pool) CreateFromPool() unsafe.Pointer {
	e := p.pool.GetFirstElement()
	if e == nil {
		return nil
	}
	p.pool.RemoveElement(e)
	p.actives.AddElement(e)
	o := (*Object)(e.GetValue())
	o.isActive = true
	return o.GetData()
}

// Sweep remove non active objects from active list
func (p *Pool) Sweep() {
	for ite := p.actives.GetIterator(); ite.HasNext(); {
		elem := ite.Next()
		o := (*Object)(elem.GetValue())
		if o.isActive == false {
			p.actives.RemoveElement(elem)
			p.pool.AddElement(elem)
		}
	}
}

// Clean deactivate all items
func (p *Pool) Clean() {
	for ite := p.actives.GetIterator(); ite.HasNext(); {
		elem := ite.Next()
		o := (*Object)(elem.GetValue())
		o.isActive = false
	}
	p.Sweep()
}
