package weapon

func Normal(factory ShotFactory, playSound bool) Weapon {
	w := &normal{baseWeapon{}}
	w.shotFactory = factory
	w.playSound = playSound
	return w
}

func Machinegun(factory ShotFactory, playSound bool) Weapon {
	w := &machingun{baseWeapon{}}
	w.shotFactory = factory
	w.playSound = playSound
	return w
}
