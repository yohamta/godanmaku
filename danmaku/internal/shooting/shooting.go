package shooting

import (
	"math/rand"
	"time"
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/quad"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"

	"github.com/yohamta/godanmaku/danmaku/internal/sound"

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
	maxEnemyShot  = 100
	maxEnemy      = 50
	maxEffects    = 100

	quadTreeDepth = 3
)

type state int

const (
	stateLoading state = iota
	statePlaying
	stateLose
	stateWin
)

type enemyPop struct {
	x, y float64
}

// Shooting represents shooting scene
type Shooting struct {
	player     *shooter.Player
	state      state
	input      *inputs.Input
	field      *field.Field
	viewCenter struct{ x, y float64 }
	enemyQueue *list.List
	tmpEnemy   *shooter.Enemy
	endTime    time.Time
	killNum    int

	// quadtree
	playersQuadTree *quad.Quad
	enemyQuadTree   *quad.Quad
	pShotQuadTree   *quad.Quad
	eShotQuadTree   *quad.Quad
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
	s.input = inputs.NewInput()

	s.field = field.NewField(ui.GetScreenWidth(), ui.GetScreenHeight())
	s.enemyQueue = list.NewList()
	s.tmpEnemy = shooter.NewEnemy(s.field, shared.EnemyShots)
	s.killNum = 0

	if shared.HealthBar == nil {
		shared.HealthBar = ui.NewHealthBar()
	}

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

	// quad tree
	x1 := s.field.GetLeft()
	x2 := s.field.GetRight()
	y1 := s.field.GetTop()
	y2 := s.field.GetBottom()
	s.playersQuadTree = quad.NewQuad(x1, x2, y1, y2, quadTreeDepth)
	s.enemyQuadTree = quad.NewQuad(x1, x2, y1, y2, quadTreeDepth)
	s.pShotQuadTree = quad.NewQuad(x1, x2, y1, y2, quadTreeDepth)
	s.eShotQuadTree = quad.NewQuad(x1, x2, y1, y2, quadTreeDepth)
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
		e := (*shooter.Enemy)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			quad.RemoveNodeFromQuad(e.GetQuadNode())
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

	switch s.state {
	case statePlaying:
		s.input.Draw(screen)
	case stateLose:
		fallthrough
	case stateWin:
		s.drawResult(screen)
	}
}

func (s *Shooting) drawResult(screen *ebiten.Image) {
	x, y := ui.GetCenterOfScreen()
	if s.state == stateLose {
		sprite.Result.SetIndex(0)
	} else {
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
	popInfo := (*enemyPop)(element.GetValue())

	enemy := (*shooter.Enemy)(shared.Enemies.CreateFromPool())
	if enemy == nil {
		return
	}
	enemy.Init(popInfo.x, popInfo.y)
	enemy.SetTarget(s.player)
}

func (s *Shooting) initEnemies() {
	enemyCount := 40

	wait := int(rand.Float64() * 10)
	radius := 300.
	for i := 0; i < enemyCount; i++ {
		// get enemy size
		s.tmpEnemy.Init(0, 0)
		x, y := s.field.GetRandamPosition(s.player.GetX(), s.player.GetY(), radius)
		s.enemyQueue.AddValue(unsafe.Pointer(&enemyPop{x: x, y: y}))

		// craete jump effect
		effect.CreateJump(x, y, wait, s.popNextEnemy)
		wait += int(rand.Float64() * 20)
	}
}

func (s *Shooting) updateQuadTree() {
	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		s.pShotQuadTree.AddNode(p.GetQuadNode())
	}

	// enemy shots
	for ite := shared.EnemyShots.GetIterator(); ite.HasNext(); {
		e := (*shot.Shot)(ite.Next().GetData())
		s.eShotQuadTree.AddNode(e.GetQuadNode())
	}

	// enemies
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		e := (*shooter.Enemy)(ite.Next().GetData())
		s.enemyQuadTree.AddNode(e.GetQuadNode())
	}

	// player
	s.playersQuadTree.AddNode(s.player.GetQuadNode())
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
			if e.IsDead() {
				s.killNum++
			}
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
