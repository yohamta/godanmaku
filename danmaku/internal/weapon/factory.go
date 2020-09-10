package weapon

// Normal creates normal wewapon
func Normal(factory shotFactory) Weapon {
	w := &normal{baseWeapon{}}
	w.shotFactory = factory
	return w
}
