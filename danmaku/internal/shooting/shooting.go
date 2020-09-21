package shooting

import (
	"math"
	"math/rand"
	"time"
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/collision"

	"github.com/yohamta/godanmaku/danmaku/internal/quad"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"

	"github.com/yohamta/godanmaku/danmaku/internal/sound"

	"github.com/yohamta/godanmaku/danmaku/internal/list"
	"github.com/yohamta/godanmaku/danmaku/internal/touch"
	"github.com/yohamta/godanmaku/danmaku/internal/ui"

	"github.com/yohamta/godanmaku/danmaku/internal/effect"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/field"

	"github.com/yohamta/godanmaku/danmaku/internal/shooter"
	"github.com/yohamta/godanmaku/danmaku/internal/shot"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/inputs"
)

const (
	maxPlayerShot  = 80
	maxEnemyShot   = 300
	maxEnemy       = 50
	maxEffects     = 100
	maxBackEffects = 100

	quadTreeDepth = 3
)

type state int

const (
	stateLoading state = iota
	statePlaying
	stateLose
	stateWin
)

type enemyData struct {
	x, y float64
}

// Shooting represents shooting scene
type Shooting struct {
	player     *shooter.Shooter
	state      state
	field      *field.Field
	viewCenter struct{ x, y float64 }
	enemyQueue *list.List
	tmpShooter *shooter.Shooter
	endTime    time.Time
	killNum    int

	// quadtree
	pShotQuadTree *quad.Quad
	eShotQuadTree *quad.Quad
}

// NewShooting returns new Shooting struct
func NewShooting() *Shooting {
	s := &Shooting{}

	s.viewCenter.x = float64(ui.GetScreenWidth() / 2)
	s.viewCenter.y = float64(ui.GetScreenHeight() / 2)

	if touch.IsTouchPrimaryInput() {
		s.viewCenter.y -= 40
	}

	s.init()
	s.setupStage()

	return s
}

func (s *Shooting) init() {
	s.state = stateLoading

	rand.Seed(time.Now().Unix())

	s.field = field.NewField()
	s.enemyQueue = list.NewList()
	s.tmpShooter = shooter.NewShooter()
	s.killNum = 0

	shared.HealthBar = ui.NewHealthBar()
	shared.GameInput = inputs.NewInput()

	// enemies
	for i := 0; i < maxEnemy; i++ {
		ptr := shooter.NewShooter()
		shared.Enemies.AddToPool(unsafe.Pointer(ptr))
	}

	// shots
	for i := 0; i < maxPlayerShot; i++ {
		ptr := shot.NewShot(s.field)
		ptr.SetQuadNode(quad.NewNode(unsafe.Pointer(ptr)))
		shared.PlayerShots.AddToPool(unsafe.Pointer(ptr))
	}

	// enemyShots
	for i := 0; i < maxEnemyShot; i++ {
		ptr := shot.NewShot(s.field)
		ptr.SetQuadNode(quad.NewNode(unsafe.Pointer(ptr)))
		shared.EnemyShots.AddToPool(unsafe.Pointer(ptr))
	}

	// effects
	for i := 0; i < maxEffects; i++ {
		shared.Effects.AddToPool(unsafe.Pointer(effect.NewEffect()))
	}

	// effects
	for i := 0; i < maxBackEffects; i++ {
		shared.BackEffects.AddToPool(unsafe.Pointer(effect.NewEffect()))
	}

	// player
	s.player = shooter.NewShooter()

	// quad tree
	x1 := s.field.GetLeft()
	x2 := s.field.GetRight()
	y1 := s.field.GetTop()
	y2 := s.field.GetBottom()
	s.pShotQuadTree = quad.NewQuad(x1, x2, y1, y2, quadTreeDepth)
	s.eShotQuadTree = quad.NewQuad(x1, x2, y1, y2, quadTreeDepth)
}

func (s *Shooting) setupStage() {
	// cleaning
	shared.Enemies.Clean()
	shared.PlayerShots.Clean()
	shared.EnemyShots.Clean()
	shared.Effects.Clean()
	shared.BackEffects.Clean()

	// player
	shooter.BuildShooter(shooter.P_ROBO1, s.player, s.field,
		s.field.GetCenterX()/2, s.field.GetCenterY()/2)

	// enemies
	s.initEnemies()

	s.state = statePlaying

	// play sound
	sound.PlayBgm(sound.BgmKindBattle)
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
	shared.GameInput.Update()

	s.updateQuadTree()
	s.checkCollision()

	player := s.player

	// player
	if player.IsDead() == false {
		player.Update()
		if shared.GameInput.Fire {
			player.Fire()
		}
	}

	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		p := (*shot.Shot)(obj.GetData())
		if p.IsActive() == false {
			obj.SetInactive()
			quad.RemoveNodeFromQuad(p.GetQuadNode())
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
			quad.RemoveNodeFromQuad(e.GetQuadNode())
			continue
		}
		e.Update()
	}

	// enemies
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shooter.Shooter)(obj.GetData())
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

	// back effects
	for ite := shared.BackEffects.GetIterator(); ite.HasNext(); {
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
	shared.BackEffects.Sweep()

	switch s.state {
	case statePlaying:
		s.checkResult()
	case stateLose:
		fallthrough
	case stateWin:
		if time.Since(s.endTime).Seconds() > 3 {
			s.setupStage()
		}
	}
}

func (s *Shooting) checkResult() {
	if shared.Enemies.GetActiveNum() == 0 && s.killNum > 0 &&
		s.enemyQueue.Length() == 0 {
		s.endTime = time.Now()
		s.state = stateWin
		return
	}

	if s.player.IsDead() {
		s.endTime = time.Now()
		s.state = stateLose
	}
}

// Draw draws the scene
func (s *Shooting) Draw(screen *ebiten.Image) {
	// update offset
	shared.OffsetX = s.player.GetX() - s.viewCenter.x
	shared.OffsetY = s.player.GetY() - s.viewCenter.y

	// draw background
	s.drawBackground(screen)

	// draw field
	s.field.Draw(screen)

	// back effects
	for ite := shared.BackEffects.GetIterator(); ite.HasNext(); {
		e := (*effect.Effect)(ite.Next().GetData())
		e.Draw(screen)
	}

	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		p.Draw(screen)
	}

	// enemies
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shooter.Shooter)(obj.GetData())
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

	switch s.state {
	case statePlaying:
		shared.GameInput.Draw(screen)
	case stateLose:
		fallthrough
	case stateWin:
		s.drawResult(screen)
	}
}

func (s *Shooting) drawBackground(screen *ebiten.Image) {
	w, h := sprite.Background.Size()
	screenW := float64(ui.GetScreenWidth())
	screenH := float64(ui.GetScreenHeight())
	centerX := screenW / 2
	centerY := screenH / 2
	scaleH := (screenW / float64(h))
	scaleW := (screenH / float64(w))
	sprite.Background.SetPosition(centerX, centerY)
	sprite.Background.DrawWithScale(screen, math.Max(scaleH, scaleW))
}

func (s *Shooting) drawResult(screen *ebiten.Image) {
	x, y := ui.GetCenterOfScreen()
	if s.state == stateLose {
		sprite.Result.SetIndex(0)
	} else {
		y -= 100
		sprite.Result.SetIndex(1)
	}
	sprite.Result.SetPosition(float64(x), float64(y))
	sprite.Result.Draw(screen)
}

func (s *Shooting) popNextEnemy() {
	q := s.enemyQueue
	if q.Length() <= 0 {
		return
	}
	element := q.GetFirstElement()
	q.RemoveElement(element)
	popInfo := (*enemyData)(element.GetValue())

	enemy := (*shooter.Shooter)(shared.Enemies.CreateFromPool())
	if enemy == nil {
		return
	}
	shooter.BuildShooter(shooter.E_ROBO1, enemy, s.field, popInfo.x, popInfo.y)
	enemy.SetTarget(s.player)
}

func (s *Shooting) initEnemies() {
	enemyCount := 30

	wait := int(rand.Float64() * 10)
	radius := 300.
	for i := 0; i < enemyCount; i++ {
		// get enemy size
		shooter.BuildShooter(shooter.E_ROBO1, s.tmpShooter, s.field, 0, 0)
		x, y := s.field.GetRandamPosition(s.player.GetX(), s.player.GetY(), radius)
		s.enemyQueue.AddValue(unsafe.Pointer(&enemyData{x: x, y: y}))

		// craete jump effect
		effect.CreateJump(x, y, wait, s.popNextEnemy)
		wait += int(rand.Float64() * 20)
	}
}

func (s *Shooting) updateQuadTree() {
	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		s.pShotQuadTree.AddNode(p, p.GetQuadNode())
	}

	// enemy shots
	for ite := shared.EnemyShots.GetIterator(); ite.HasNext(); {
		e := (*shot.Shot)(ite.Next().GetData())
		s.eShotQuadTree.AddNode(e, e.GetQuadNode())
	}
}

func (s *Shooting) checkCollision() {
	// player shots
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		enemy := (*shooter.Shooter)(ite.Next().GetData())
		if enemy.IsDead() {
			continue
		}
		qd := s.pShotQuadTree.SearchQuad(enemy)
		for ite2 := qd.GetIterator(); ite2.HasNext(); {
			shot := (*shot.Shot)(ite2.Next().GetItem())
			if shot.IsActive() == false {
				continue
			}
			if collision.IsCollideWith(enemy, shot) == false {
				continue
			}
			enemy.AddDamage(1)
			shot.OnHit()
			if enemy.IsDead() {
				s.killNum++
			}
		}
	}

	// enemy shots
	{
		qd := s.eShotQuadTree.SearchQuad(s.player)
		for ite2 := qd.GetIterator(); ite2.HasNext(); {
			shot := (*shot.Shot)(ite2.Next().GetItem())
			if s.player.IsDead() {
				break
			}
			if shot.IsActive() == false {
				continue
			}
			if collision.IsCollideWith(s.player, shot) == false {
				continue
			}
			s.player.AddDamage(1)
			shot.OnHit()
		}
	}
}
