package effects

// Effect represents the base of player, enemy, shots
type Effect struct {
	x        float64
	y        float64
	isActive bool
}

// IsActive returns if this is active
func (e *Effect) IsActive() bool {
	return e.isActive
}
