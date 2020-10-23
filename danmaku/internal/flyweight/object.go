package flyweight

import (
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/linkedlist"
)

type Object struct {
	data     unsafe.Pointer
	isActive bool
	elem     *linkedlist.Element
}

func (o *Object) GetData() unsafe.Pointer {
	return o.data
}

func (o *Object) SetInactive() {
	o.isActive = false
}
