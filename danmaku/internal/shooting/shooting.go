package shooting

import (
	"math/rand"
	"time"
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/list"
	"github.com/yohamta/godanmaku/danmaku/internal/touch"
	"github.com/yohamta/godanmaku/danmaku/internal/ui"

	"github.com/yohamta/godanmaku/danmaku/internal/effect"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/util"

	"github.com/yohamta/godanmaku/danmaku/internal/shooter"
	"github.com/yohamta/godanmaku/danmaku/internal/shot"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/inputs"
)

const (
	maxPlayerShot = 80
	maxEnemyShot  = 70
	maxEnemy      = 50
	maxEffects    = 100
)

type gameState int

const (
	gameStateLoading gameState = iota
	gameStatePlaying
)

type enemyPop struct {
	x, y float64
}

// Shooting represents shooting scene
type Shooting struct {
	player     *shooter.Player
	state      gameState
	input      *inputs.Input
	field      *field.Field
	viewCenter struct{ x, y float64 }
	enemyQueue *list.List
	tmpEnemy   *shooter.Enemy
}

// NewShooting returns new Shooting struct
func NewShooting() *Shooting {
	s := &Shooting{}

	s.viewCenter.x = float64(ui.GetScreenWidth() / 2)
	s.viewCenter.y = float64(ui.GetScreenHeight() / 2)

	if touch.IsTouchPrimaryInput() {
		s.viewCenter.y -= 40
	}

	s.state = gameStateLoading
	s.init()
	s.setupStage()
	s.state = gameStatePlaying

	return s
}

func (s *Shooting) init() {
	rand.Seed(time.Now().Unix())
	s.input = inputs.NewInput()

	s.field = field.NewField(ui.GetScreenWidth(), ui.GetScreenHeight())
	s.enemyQueue = list.NewList()
	s.tmpEnemy = shooter.NewEnemy(s.field, shared.EnemyShots)

	// enemies
	for i := 0; i < maxEnemy; i++ {
		shared.Enemies.AddToPool(unsafe.Pointer(shooter.NewEnemy(s.field, shared.EnemyShots)))
	}

	// shots
	for i := 0; i < maxPlayerShot; i++ {
		shared.PlayerShots.AddToPool(unsafe.Pointer(shot.NewShot(s.field)))
	}

	// enemyShots
	for i := 0; i < maxEnemyShot; i++ {
		shared.EnemyShots.AddToPool(unsafe.Pointer(shot.NewShot(s.field)))
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
	s.player = shooter.NewPlayer(s.field, shared.PlayerShots)
	s.player.Init()

	// enemies
	s.initEnemies()
}

// GetPosition returns view position
func (s *Shooting) GetPosition() (int, int) {
	return 0, 0
}

// GetSize returns view size
func (s *Shooting) GetSize() (int, int) {
	return ui.GetScreenWidth(), ui.GetScreenHeight()
}

// Update updates the scene
func (s *Shooting) Update() {
	s.input.Update()

	s.checkCollision()

	player := s.player
	input := s.input

	// player
	if player.IsDead() == false {
		player.Update(input.Horizontal, input.Vertical, input.Fire)
		if input.Fire {
			player.Fire()
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
		p.Update()
	}

	// enemy shots
	for ite := shared.EnemyShots.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shot.Shot)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			continue
		}
		e.Update()
	}

	// enemies
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shooter.Enemy)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			continue
		}
		e.Update()
		if player.IsDead() == false {
			e.Fire()
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
	// update offset
	shared.OffsetX = s.player.GetX() - s.viewCenter.x
	shared.OffsetY = s.player.GetY() - s.viewCenter.y

	// draw background
	s.field.Draw(screen)

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

	s.input.Draw(screen)
}

func (s *Shooting) popNextEnemy() {
	q := s.enemyQueue
	if q.Length() <= 0 {
		return
	}
	element := q.GetFirstElement()
	q.RemoveElement(element)
	popInfo := (*enemyPop)(element.GetValue())

	enemy := (*shooter.Enemy)(shared.Enemies.CreateFromPool())
	if enemy == nil {
		return
	}
	enemy.Init(popInfo.x, popInfo.y)
	enemy.SetTarget(s.player)
}

func (s *Shooting) initEnemies() {
	enemyCount := 30

	wait := int(rand.Float64() * 10)
	for i := 0; i < enemyCount; i++ {
		// get enemy size
		s.tmpEnemy.Init(0, 0)
		x, y := s.field.GetRandamPosition(s.player.GetX(), s.player.GetY(), 200)
		s.enemyQueue.AddValue(unsafe.Pointer(&enemyPop{x: x, y: y}))

		// craete jump effect
		effect.CreateJump(x, y, wait, s.popNextEnemy)
		wait += int(rand.Float64() * 30)
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
