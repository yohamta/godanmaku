package shooting

import (
	"image/color"
	"math/rand"
	"time"
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/effect"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
	"github.com/yohamta/godanmaku/danmaku/internal/weapon"

	"github.com/yohamta/godanmaku/danmaku/internal/shooter"
	"github.com/yohamta/godanmaku/danmaku/internal/shot"

	"github.com/yohamta/godanmaku/danmaku/internal/ui"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/inputs"
)

const (
	maxPlayerShot = 80
	maxEnemyShot  = 70
	maxEnemy      = 50
	maxEffects    = 100
)

var (
	backgroundColor = color.RGBA{0x00, 0x00, 0x00, 0xff}
)

type gameState int

const (
	gameStateLoading gameState = iota
	gameStatePlaying
)

// Shooting represents shooting scene
type Shooting struct {
	screenWidth  int
	screenHeight int
	player       *shooter.Player
	state        gameState
	background   *ui.Box
	input        *inputs.Input
	fld          *field.Field
}

// NewShooting returns new Shooting struct
func NewShooting(screenWidth, screenHeight int) *Shooting {
	s := &Shooting{}

	s.screenWidth = screenWidth
	s.screenHeight = screenHeight

	s.state = gameStateLoading
	s.init()
	s.setupStage()
	s.state = gameStatePlaying

	return s
}

func (s *Shooting) init() {
	rand.Seed(time.Now().Unix())
	s.input = inputs.NewInput(s.screenWidth, s.screenHeight)

	fld := field.NewField()
	s.fld = fld

	s.background = ui.NewBox(0, int(fld.GetBottom()),
		s.screenWidth,
		s.screenHeight-int(fld.GetBottom()-fld.GetTop()),
		backgroundColor)

	// enemies
	for i := 0; i < maxEnemy; i++ {
		shared.Enemies.AddToPool(unsafe.Pointer(shooter.NewEnemy(fld, shared.EnemyShots)))
	}

	// shots
	for i := 0; i < maxPlayerShot; i++ {
		shared.PlayerShots.AddToPool(unsafe.Pointer(shot.NewShot(fld)))
	}

	// enemyShots
	for i := 0; i < maxEnemyShot; i++ {
		shared.EnemyShots.AddToPool(unsafe.Pointer(shot.NewShot(fld)))
	}

	// effects
	for i := 0; i < maxEffects; i++ {
		shared.Effects.AddToPool(unsafe.Pointer(effect.NewEffect()))
	}
}

func (s *Shooting) setupStage() {
	// cleaning
	shared.Enemies.Clean()
	shared.PlayerShots.Clean()
	shared.EnemyShots.Clean()
	shared.Effects.Clean()

	// player
	s.player = shooter.NewPlayer(s.fld, shared.PlayerShots)
	s.player.Init()
	s.player.SetMainWeapon(weapon.NewNormal(shot.KindPlayerNormal))
	s.player.SetField(s.fld)

	// enemies
	s.initEnemies()
}

// Update updates the scene
func (s *Shooting) Update() {
	s.input.Update()

	s.checkCollision()

	player := s.player
	input := s.input

	// player
	if player.IsDead() == false {
		player.Move(input.Horizontal, input.Vertical, input.Fire)
		if input.Fire {
			player.FireWeapon()
		}
	}

	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		p := (*shot.Shot)(obj.GetData())
		if p.IsActive() == false {
			obj.SetInactive()
			continue
		}
		p.Move()
	}

	// enemy shots
	for ite := shared.EnemyShots.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shot.Shot)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			continue
		}
		e.Move()
	}

	// enemies
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shooter.Enemy)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			continue
		}
		e.Move()
		if player.IsDead() == false {
			e.FireWeapon()
		}
	}

	// effects
	for ite := shared.Effects.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*effect.Effect)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			continue
		}
		e.Update()
	}

	shared.EnemyShots.Sweep()
	shared.PlayerShots.Sweep()
	shared.Enemies.Sweep()
	shared.Effects.Sweep()

	if player.IsDead() && shared.Effects.GetActiveNum() == 0 {
		s.setupStage()
	}
}

// Draw draws the scene
func (s *Shooting) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})

	s.fld.Draw(screen)

	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		p.Draw(screen)
	}

	// enemies
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shooter.Enemy)(obj.GetData())
		e.Draw(screen)
	}

	if s.player.IsDead() == false {
		s.player.Draw(screen)
	}

	// enemy shots
	for ite := shared.EnemyShots.GetIterator(); ite.HasNext(); {
		e := (*shot.Shot)(ite.Next().GetData())
		e.Draw(screen)
	}

	// effects
	for ite := shared.Effects.GetIterator(); ite.HasNext(); {
		e := (*effect.Effect)(ite.Next().GetData())
		e.Draw(screen)
	}

	s.background.Draw(screen)
	s.input.Draw(screen)
}

func (s *Shooting) initEnemies() {
	enemyCount := 20

	for i := 0; i < enemyCount; i++ {
		enemy := (*shooter.Enemy)(shared.Enemies.CreateFromPool())
		if enemy == nil {
			return
		}
		enemy.Init()
		enemy.SetTarget(s.player)
	}
}

func (s *Shooting) checkCollision() {
	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		for ite2 := shared.Enemies.GetIterator(); ite2.HasNext(); {
			e := (*shooter.Enemy)(ite2.Next().GetData())
			if util.IsCollideWith(e, p) == false {
				continue
			}
			e.AddDamage(1)
			p.OnHit()
		}
	}

	// enemy shots
	if s.player.IsDead() == false {
		for ite := shared.EnemyShots.GetIterator(); ite.HasNext(); {
			e := (*shot.Shot)(ite.Next().GetData())
			if util.IsCollideWith(s.player, e) == false {
				continue
			}
			s.player.AddDamage(1)
			e.OnHit()
		}
	}
}
