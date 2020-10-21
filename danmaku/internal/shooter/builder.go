package shooter

import (
	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/shot"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
	"github.com/yohamta/godanmaku/danmaku/internal/weapon"
)

func BuildFunnel(funnel *Shooter, owner *Shooter, f *field.Field) {
	BuildShooter(P_FUNNEL, funnel, f, owner.GetX(), owner.GetY())
	funnel.owner = owner
	owner.funnel = append(owner.funnel, funnel)
	funnel.Update()
}

func BuildShooter(kind kind, sh *Shooter, f *field.Field, x, y float64) {
	sh.isActive = true
	sh.field = f
	sh.SetPosition(x, y)
	sh.funnel = nil
	sh.owner = nil

	switch kind {
	case P_ROBO1:
		sh.setSize(16, 16)
		sh.SetSpeed(0.96, 90)
		sh.SetSpeed(2, 270)
		sh.SetWeapon(weapon.Normal(shot.PlayerShot, true))

		sh.collisionBox = collision.GetCollisionBox("P_ROBO_1")
		sh.life = 10
		sh.maxLife = sh.life
		sh.spr = sprite.Player
		sh.controller = player

		break
	case P_FUNNEL:
		sh.setSize(10, 10)
		sh.SetWeapon(weapon.Machinegun(shot.BlueLaser, true))

		sh.collisionBox = collision.GetCollisionBox("NULL")
		sh.spr = sprite.Funnel
		sh.controller = funnel

		break
	case E_ROBO1:
		sh.setSize(24, 24)
		sh.SetSpeed(0.96, 90)
		sh.SetWeapon(weapon.Normal(shot.EnemyShot, false))

		sh.collisionBox = collision.GetCollisionBox("E_ROBO1")
		sh.life = 3
		sh.maxLife = sh.life
		sh.spr = sprite.Enemy1
		sh.controller = npc
		break
	}
	sh.init()
}
