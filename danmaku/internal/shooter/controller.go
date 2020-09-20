package shooter

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/shot"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/weapon"
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

// BuildShooter builds shooter
func BuildShooter(kind kind, sh *Shooter, f *field.Field, x, y float64) {
	sh.isActive = true
	sh.field = f
	sh.SetPosition(x, y)

	switch kind {
	case P_ROBO1:
		sh.setSize(16, 16)
		sh.SetSpeed(0.96, 90)
		sh.SetSpeed(2, 270)
		sh.SetWeapon(weapon.Machinegun(shot.BlueLaser, true))

		sh.collisionBox = collision.GetCollisionBox("P_ROBO_1")
		sh.life = 10
		sh.maxLife = sh.life
		sh.spr = sprite.Player
		sh.controller = player

		break
	case E_ROBO1:
		sh.setSize(24, 24)
		sh.SetSpeed(0.96, 90)
		sh.SetWeapon(weapon.Normal(shot.EnemyShot, false))

		sh.collisionBox = collision.GetCollisionBox("E_ROBO1")
		sh.life = 3
		sh.maxLife = sh.life
		sh.spr = sprite.Enemy1
		sh.controller = enemy
		break
	}
	sh.init()
}
