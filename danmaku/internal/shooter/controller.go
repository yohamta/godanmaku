package shooter

import (
	"github.com/hajimehoshi/ebiten"
)

type controller interface {
	init(sh *Shooter)
	update(sh *Shooter)
	draw(sh *Shooter, screen *ebiten.Image)
}

type kind int

const (
	P_ROBO1 kind = iota
	E_ROBO1
)

var (
	enemy  = new(enemyController)
	player = new(playerController)
)
