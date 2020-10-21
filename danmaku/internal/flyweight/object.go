package flyweight

import (
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/list"
)

type Object struct {
	data     unsafe.Pointer
	isActive bool
	elem     *list.Element
}

func (o *Object) GetData() unsafe.Pointer {
	return o.data
}

func (o *Object) SetInactive() {
	o.isActive = false
}
