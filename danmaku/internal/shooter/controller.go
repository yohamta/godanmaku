package shooter

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Controller interface {
	init(sh *Shooter)
	update(sh *Shooter)
	draw(sh *Shooter, screen *ebiten.Image)
}

type kind int

const (
	P_ROBO1 kind = iota
	P_FUNNEL
	E_ROBO1
	GRAZE
)

var (
	npc    = new(NPCController)
	player = new(PlayerController)
	funnel = new(FunnelController)
	graze  = new(GrazeController)
)
