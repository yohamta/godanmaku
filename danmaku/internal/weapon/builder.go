package weapon

// Normal creates normal wewapon
func Normal(factory shotFactory, playSound bool) Weapon {
	w := &normal{baseWeapon{}}
	w.shotFactory = factory
	w.playSound = playSound
	return w
}
