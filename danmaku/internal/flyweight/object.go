package flyweight

import (
	"unsafe"

	"github.com/yotahamada/godanmaku/danmaku/internal/list"
)

// Object represents object
type Object struct {
	data     unsafe.Pointer
	isActive bool
	elem     *list.Element
}

// GetData returns data
func (o *Object) GetData() unsafe.Pointer {
	return o.data
}

// SetInactive marks this object inactive
func (o *Object) SetInactive() {
	o.isActive = false
}
