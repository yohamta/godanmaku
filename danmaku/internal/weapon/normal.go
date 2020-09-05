package weapon

import (
	"time"
)

// Normal represents player's weapon
type Normal struct{ baseWeapon }

// NewNormal creates new struct
func NewNormal(factory shotFactoryFunction) *Normal {
	w := &Normal{baseWeapon{}}
	w.shotFactory = factory

	return w
}

// Fire create shots
func (w *Normal) Fire(x, y float64, degree int) {
	if time.Since(w.lastShotTime).Milliseconds() < 350 {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, degree)
}
