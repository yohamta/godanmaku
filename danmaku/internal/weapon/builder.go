package weapon

// Normal creates normal wewapon
func Normal(factory shotFactory, playSound bool) Weapon {
	w := &normal{baseWeapon{}}
	w.shotFactory = factory
	w.playSound = playSound
	return w
}

// Machinegun creates normal wewapon
func Machinegun(factory shotFactory, playSound bool) Weapon {
	w := &machingun{baseWeapon{}}
	w.shotFactory = factory
	w.playSound = playSound
	return w
}
