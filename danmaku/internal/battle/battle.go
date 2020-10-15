package battle

import (
	"image/color"
	"math"
	"math/rand"
	"time"
	"unsafe"

	"github.com/yotahamada/furex"

	"github.com/yotahamada/godanmaku/danmaku/internal/collision"
	"github.com/yotahamada/godanmaku/danmaku/internal/field"
	"github.com/yotahamada/godanmaku/danmaku/internal/paint"

	"github.com/yotahamada/godanmaku/danmaku/internal/quad"
	"github.com/yotahamada/godanmaku/danmaku/internal/sprite"

	"github.com/yotahamada/godanmaku/danmaku/internal/sound"

	"github.com/yotahamada/godanmaku/danmaku/internal/list"
	"github.com/yotahamada/godanmaku/danmaku/internal/touch"
	"github.com/yotahamada/godanmaku/danmaku/internal/ui"

	"github.com/yotahamada/godanmaku/danmaku/internal/effect"
	"github.com/yotahamada/godanmaku/danmaku/internal/shared"

	"github.com/yotahamada/godanmaku/danmaku/internal/shooter"
	"github.com/yotahamada/godanmaku/danmaku/internal/shot"

	"github.com/hajimehoshi/ebiten"
	"github.com/yotahamada/godanmaku/danmaku/internal/inputs"
)

const (
	maxPlayerShot  = 80
	maxEnemyShot   = 300
	maxEnemy       = 50
	maxEffects     = 100
	maxBackEffects = 100
	quadTreeDepth  = 3
)

type BattleState int

const (
	stateLoading BattleState = iota
	statePlaying
	stateLose
	stateWin
)

type EnemyData struct {
	x, y float64
}

type Battle struct {
	ui *furex.Controller

	player      *shooter.Shooter
	state       BattleState
	field       *field.Field
	center      struct{ x, y float64 }
	enemyQueue  *list.List
	tmpShooter  *shooter.Shooter
	endTime     time.Time
	killNum     int
	updateCount int

	dispTextTime time.Time
	dispText     string

	pShotQuadTree *quad.Quad
	eShotQuadTree *quad.Quad
}

func NewBattle() *Battle {
	b := &Battle{}

	if touch.IsTouchPrimaryInput() {
		b.center.y -= 40
	}

	b.initBattle()
	b.initStage()
	b.initUI()

	return b
}

func (b *Battle) Layout(width, height int) {
	b.ui.Layout(0, 0, width, height)
	b.center.x = float64(width / 2)
	b.center.y = float64(height / 2)
}

func (b *Battle) Update() {
	b.updateCount++
	shared.GameInput.Update()

	b.updateQuadTree()
	b.checkCollision()

	player := b.player

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

	switch b.state {
	case statePlaying:
		b.checkResult()
	case stateLose:
		fallthrough
	case stateWin:
		if time.Since(b.endTime).Seconds() > 3 {
			b.initStage()
		}
	}
}

func (b *Battle) Draw(screen *ebiten.Image) {
	// update offset
	shared.OffsetX = b.player.GetX() - b.center.x
	shared.OffsetY = b.player.GetY() - b.center.y

	// draw background
	b.drawBackground(screen)

	// draw field
	b.field.Draw(screen)

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

	if b.player.IsDead() == false {
		b.player.Draw(screen)
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

	if time.Since(b.dispTextTime).Seconds() <= 1 {
		w := 3 * 24
		h := 24
		shouldPaint := true
		if time.Since(b.dispTextTime).Seconds() < 0.3 {
			if b.updateCount%6 > 3 {
				shouldPaint = false
			}
		}
		if shouldPaint {
			paint.DrawText(screen, b.dispText, ui.ScreenWidth/2-w/2,
				h+10, color.White, paint.FontSizeXLarge)
		}
	}

	switch b.state {
	case statePlaying:
		shared.GameInput.Draw(screen)
	case stateLose:
		fallthrough
	case stateWin:
		b.drawResult(screen)
	}

	b.ui.Draw(screen)
}

func (b *Battle) initBattle() {
	b.state = stateLoading

	rand.Seed(time.Now().Unix())

	b.field = field.NewField()
	b.enemyQueue = list.NewList()
	b.tmpShooter = shooter.NewShooter()
	b.killNum = 0

	shared.HealthBar = ui.NewHealthBar()
	shared.GameInput = inputs.NewInput()

	// enemies
	for i := 0; i < maxEnemy; i++ {
		ptr := shooter.NewShooter()
		shared.Enemies.AddToPool(unsafe.Pointer(ptr))
	}

	// shots
	for i := 0; i < maxPlayerShot; i++ {
		ptr := shot.NewShot(b.field)
		ptr.SetQuadNode(quad.NewNode(unsafe.Pointer(ptr)))
		shared.PlayerShots.AddToPool(unsafe.Pointer(ptr))
	}

	// enemyShots
	for i := 0; i < maxEnemyShot; i++ {
		ptr := shot.NewShot(b.field)
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
	b.player = shooter.NewShooter()

	// quad tree
	x1 := b.field.GetLeft()
	x2 := b.field.GetRight()
	y1 := b.field.GetTop()
	y2 := b.field.GetBottom()
	b.pShotQuadTree = quad.NewQuad(x1, x2, y1, y2, quadTreeDepth)
	b.eShotQuadTree = quad.NewQuad(x1, x2, y1, y2, quadTreeDepth)
}

func (b *Battle) initUI() {
	b.ui = furex.NewController()
	// TODO:
}

func (b *Battle) initStage() {
	// cleaning
	shared.Enemies.Clean()
	shared.PlayerShots.Clean()
	shared.EnemyShots.Clean()
	shared.Effects.Clean()
	shared.BackEffects.Clean()

	// player
	shooter.BuildShooter(shooter.P_ROBO1, b.player, b.field,
		b.field.GetCenterX()/2, b.field.GetCenterY()/2)

	// enemies
	b.initEnemies()

	b.state = statePlaying

	// play sound
	sound.PlayBgm(sound.BgmKindBattle)
}

func (b *Battle) checkResult() {
	if shared.Enemies.GetActiveNum() == 0 && b.killNum > 0 &&
		b.enemyQueue.Length() == 0 {
		b.endTime = time.Now()
		b.state = stateWin
		return
	}

	if b.player.IsDead() {
		b.endTime = time.Now()
		b.state = stateLose
	}
}

func (b *Battle) drawBackground(screen *ebiten.Image) {
	w, h := sprite.Background.Size()
	screenW := float64(ui.ScreenWidth)
	screenH := float64(ui.ScreenHeight)
	centerX := screenW / 2
	centerY := screenH / 2
	scaleH := (screenW / float64(h))
	scaleW := (screenH / float64(w))
	sprite.Background.SetPosition(centerX, centerY)
	sprite.Background.DrawWithScale(screen, math.Max(scaleH, scaleW))
}

func (b *Battle) drawResult(screen *ebiten.Image) {
	x, y := ui.GetCenterOfScreen()
	if b.state == stateLose {
		sprite.Result.SetIndex(0)
	} else {
		y -= 100
		sprite.Result.SetIndex(1)
	}
	sprite.Result.SetPosition(float64(x), float64(y))
	sprite.Result.Draw(screen)
}

func (b *Battle) popNextEnemy() {
	q := b.enemyQueue
	if q.Length() <= 0 {
		return
	}
	element := q.GetFirstElement()
	q.RemoveElement(element)
	popInfo := (*EnemyData)(element.GetValue())

	enemy := (*shooter.Shooter)(shared.Enemies.CreateFromPool())
	if enemy == nil {
		return
	}
	shooter.BuildShooter(shooter.E_ROBO1, enemy, b.field, popInfo.x, popInfo.y)
	enemy.SetTarget(b.player)
}

func (b *Battle) initEnemies() {
	enemyCount := 30

	wait := int(rand.Float64() * 10)
	radius := 300.
	for i := 0; i < enemyCount; i++ {
		// get enemy size
		shooter.BuildShooter(shooter.E_ROBO1, b.tmpShooter, b.field, 0, 0)
		x, y := b.field.GetRandamPosition(b.player.GetX(), b.player.GetY(), radius)
		b.enemyQueue.AddValue(unsafe.Pointer(&EnemyData{x: x, y: y}))

		// craete jump effect
		effect.CreateJump(x, y, wait, b.popNextEnemy)
		wait += int(rand.Float64() * 20)
	}
}

func (b *Battle) updateQuadTree() {
	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		b.pShotQuadTree.AddNode(p, p.GetQuadNode())
	}

	// enemy shots
	for ite := shared.EnemyShots.GetIterator(); ite.HasNext(); {
		e := (*shot.Shot)(ite.Next().GetData())
		b.eShotQuadTree.AddNode(e, e.GetQuadNode())
	}
}

func (b *Battle) checkCollision() {
	// player shots
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		enemy := (*shooter.Shooter)(ite.Next().GetData())
		if enemy.IsDead() {
			continue
		}
		qd := b.pShotQuadTree.SearchQuad(enemy)
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
				b.killNum++
				b.dispTextTime = time.Now()
				b.dispText = "[撃破]"
			} else {
				b.dispTextTime = time.Now()
				b.dispText = "[命中]"
			}
		}
	}

	// enemy shots
	{
		qd := b.eShotQuadTree.SearchQuad(b.player)
		for ite2 := qd.GetIterator(); ite2.HasNext(); {
			shot := (*shot.Shot)(ite2.Next().GetItem())
			if b.player.IsDead() {
				break
			}
			if shot.IsActive() == false {
				continue
			}
			if collision.IsCollideWith(b.player, shot) == false {
				continue
			}
			b.player.AddDamage(1)
			shot.OnHit()

			b.dispTextTime = time.Now()
			b.dispText = "[被弾]"
		}
	}
}
