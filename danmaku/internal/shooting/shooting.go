package shooting

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
	"unsafe"

	"github.com/yohamta/furex"

	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/linkedlist"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"

	"github.com/yohamta/godanmaku/danmaku/internal/quadtree"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"

	"github.com/yohamta/godanmaku/danmaku/internal/sound"

	"github.com/yohamta/godanmaku/danmaku/internal/effect"

	"github.com/yohamta/godanmaku/danmaku/internal/shooter"
	"github.com/yohamta/godanmaku/danmaku/internal/shot"

	"github.com/hajimehoshi/ebiten"
)

const (
	maxPlayerShot   = 300
	maxPlayerFunnel = 20
	maxEnemyShot    = 300
	maxEnemy        = 50
	maxEffects      = 100
	maxBackEffects  = 100
	quadTreeDepth   = 3
)

type State int

const (
	stateLoading State = iota
	statePlaying
	stateLose
	stateWin
)

type EnemyData struct {
	x, y float64
}

var (
	isInitUIDone bool
	battleView   *furex.View

	player      *shooter.Shooter
	state       State
	fld         *field.Field
	enemyQueue  *linkedlist.List
	tmpShooter  *shooter.Shooter
	endTime     time.Time
	killNum     int
	updateCount int

	dispTextTime time.Time
	dispText     string

	pShotQuadTree *quadtree.Quadtree
	eShotQuadTree *quadtree.Quadtree

	screenSize   image.Point
	screenCenter image.Point

	fireButton *FireButton
	joystick   *Joystick
)

type Shooting struct{}

func NewShooting() *Shooting {
	s := &Shooting{}

	loadResources()

	initObjects()
	initStage()

	return s
}

func (s *Shooting) Layout(width, height int) {
	screenSize = image.Pt(width, height)
	screenCenter.X = screenSize.X / 2
	screenCenter.Y = screenSize.Y/2 - 50
	shared.ScreenSize = screenSize

	if battleView != nil {
		battleView.Layout(0, 0, screenSize.X, screenSize.Y)
	}
	if isInitUIDone == false {
		initUI()
		isInitUIDone = true
	}
}

func (s *Shooting) Update() {
	updateCount++

	updateInput()
	updateQuadTree()
	checkCollision()
	updateObjects()

	battleView.Update()

	switch state {
	case statePlaying:
		checkResult()
	case stateLose:
		fallthrough
	case stateWin:
		if time.Since(endTime).Seconds() > 3 {
			initStage()
		}
	}
}

func (s *Shooting) Draw(screen *ebiten.Image) {
	shared.OffsetX = player.GetX() - float64(screenCenter.X)
	shared.OffsetY = player.GetY() - float64(screenCenter.Y)

	drawBackground(screen)
	drawObjects(screen)
	drawMessages(screen)

	switch state {
	case statePlaying:
	case stateLose:
		fallthrough
	case stateWin:
		drawResult(screen)
	}

	battleView.Draw(screen)
}

func initObjects() {
	state = stateLoading

	rand.Seed(time.Now().Unix())

	fld = field.NewField()
	enemyQueue = linkedlist.NewList()
	tmpShooter = shooter.NewShooter()
	killNum = 0

	// enemies
	for i := 0; i < maxEnemy; i++ {
		ptr := shooter.NewShooter()
		shared.Enemies.AddToPool(unsafe.Pointer(ptr))
	}

	// shots
	for i := 0; i < maxPlayerShot; i++ {
		ptr := shot.NewShot(fld)
		ptr.SetQuadtreeNode(quadtree.NewNode(unsafe.Pointer(ptr)))
		shared.PlayerShots.AddToPool(unsafe.Pointer(ptr))
	}

	// funnels
	for i := 0; i < maxPlayerFunnel; i++ {
		ptr := shooter.NewShooter()
		shared.PlayerFunnels.AddToPool(unsafe.Pointer(ptr))
	}

	// enemyShots
	for i := 0; i < maxEnemyShot; i++ {
		ptr := shot.NewShot(fld)
		ptr.SetQuadtreeNode(quadtree.NewNode(unsafe.Pointer(ptr)))
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
	player = shooter.NewShooter()

	// quad tree
	x1 := fld.GetLeft()
	x2 := fld.GetRight()
	y1 := fld.GetTop()
	y2 := fld.GetBottom()
	pShotQuadTree = quadtree.NewQuadtree(x1, x2, y1, y2, quadTreeDepth)
	eShotQuadTree = quadtree.NewQuadtree(x1, x2, y1, y2, quadTreeDepth)
}

func initUI() {
	battleView = furex.NewView()
	battleView.Layout(0, 0, screenSize.X, screenSize.Y)

	flex := furex.NewFlex(0, 0, screenSize.X, screenSize.Y)
	battleView.AddLayer(furex.NewLayerWithContainer(flex))

	joystick = NewJoystick()
	flex.AddChild(joystick)

	fireButton = NewFireButton()
	flex.AddChild(fireButton)

	battleView.AddLayer(furex.NewLayerWithContainer(flex))
}

func initStage() {
	shared.Enemies.Clean()
	shared.PlayerShots.Clean()
	shared.PlayerFunnels.Clean()
	shared.EnemyShots.Clean()
	shared.Effects.Clean()
	shared.BackEffects.Clean()

	shooter.BuildShooter(shooter.P_ROBO1, player, fld,
		fld.GetCenterX()/2, fld.GetCenterY()/2)

	for i := 0; i < 7; i++ {
		// test
		funnel := (*shooter.Shooter)(shared.PlayerFunnels.CreateFromPool())
		if funnel == nil {
			return
		}
		shooter.BuildFunnel(funnel, player, fld)
	}

	initEnemies()

	sound.PlayBgm(sound.BgmKindBattle)

	state = statePlaying
}

func checkResult() {
	if shared.Enemies.GetActiveNum() == 0 && killNum > 0 &&
		enemyQueue.Length() == 0 {
		endTime = time.Now()
		state = stateWin
		return
	}

	if player.IsDead() {
		endTime = time.Now()
		state = stateLose
	}
}

func drawBackground(screen *ebiten.Image) {
	w, h := sprite.Background.Size()
	screenW := float64(screenSize.X)
	screenH := float64(screenSize.Y)
	centerX := screenW / 2
	centerY := screenH / 2
	scaleH := (screenW / float64(h))
	scaleW := (screenH / float64(w))
	sprite.Background.SetPosition(centerX, centerY)
	sprite.Background.DrawWithScale(screen, math.Max(scaleH, scaleW))

	fld.Draw(screen)
}

func drawResult(screen *ebiten.Image) {
	center := screenCenter
	center.Y -= 50
	if state == stateLose {
		sprite.Result.SetIndex(0)
	} else {
		sprite.Result.SetIndex(1)
	}
	sprite.Result.SetPosition(float64(center.X), float64(center.Y))
	sprite.Result.Draw(screen)
}

func popNextEnemy() {
	q := enemyQueue
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
	shooter.BuildShooter(shooter.E_ROBO1, enemy, fld, popInfo.x, popInfo.y)
	enemy.SetTarget(player)
}

func initEnemies() {
	enemyCount := 100

	wait := int(rand.Float64() * 10)
	radius := 300.
	for i := 0; i < enemyCount; i++ {
		// get enemy size
		shooter.BuildShooter(shooter.E_ROBO1, tmpShooter, fld, 0, 0)
		x, y := fld.GetRandamPosition(player.GetX(), player.GetY(), radius)
		enemyQueue.AddValue(unsafe.Pointer(&EnemyData{x: x, y: y}))

		// craete jump effect
		effect.CreateJump(x, y, wait, popNextEnemy)
		wait += int(rand.Float64() * 20)
	}
}

func updateQuadTree() {
	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		pShotQuadTree.AddNode(p, p.GetQuadtreeNode())
	}

	// enemy shots
	for ite := shared.EnemyShots.GetIterator(); ite.HasNext(); {
		e := (*shot.Shot)(ite.Next().GetData())
		eShotQuadTree.AddNode(e, e.GetQuadtreeNode())
	}
}

func checkCollision() {
	// player shots
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		enemy := (*shooter.Shooter)(ite.Next().GetData())
		if enemy.IsDead() {
			continue
		}
		qd := pShotQuadTree.SearchQuadtree(enemy)
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
				killNum++
				dispTextTime = time.Now()
				dispText = "[撃破]"
			} else {
				dispTextTime = time.Now()
				dispText = "[命中]"
			}
		}
	}

	// enemy shots
	{
		qd := eShotQuadTree.SearchQuadtree(player)
		for ite2 := qd.GetIterator(); ite2.HasNext(); {
			shot := (*shot.Shot)(ite2.Next().GetItem())
			if player.IsDead() {
				break
			}
			if shot.IsActive() == false {
				continue
			}
			if collision.IsCollideWith(player, shot) == false {
				continue
			}
			player.AddDamage(1)
			shot.OnHit()

			dispTextTime = time.Now()
			dispText = "[被弾]"
		}
	}
}

func loadResources() {
	paint.LoadFonts()
	sprite.LoadSprites()
	sound.Load()
}

func updateInput() {
	v := 0.
	h := 0.

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		v = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		v = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		h = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		h = -1
	}

	if v == 0 && h == 0 && joystick.isReadingTouch {
		v = joystick.vertical
		h = joystick.horizontal
	}

	shared.GameInput.Vertical = v
	shared.GameInput.Horizontal = h
	shared.GameInput.Fire = fireButton.isPressing || ebiten.IsKeyPressed(ebiten.KeySpace)
}

func updateObjects() {
	// player
	if player.IsDead() == false {
		player.Update()
		if shared.GameInput.Fire {
			player.Fire()
		}
	}

	// player funnels
	for ite := shared.PlayerFunnels.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		f := (*shooter.Shooter)(obj.GetData())
		if f.IsActive() == false {
			obj.SetInactive()
			continue
		}
		f.Update()
	}

	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		p := (*shot.Shot)(obj.GetData())
		if p.IsActive() == false {
			obj.SetInactive()
			quadtree.RemoveNodeFromQuadtree(p.GetQuadtreeNode())
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
			quadtree.RemoveNodeFromQuadtree(e.GetQuadtreeNode())
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
	shared.PlayerFunnels.Sweep()
	shared.Enemies.Sweep()
	shared.Effects.Sweep()
	shared.BackEffects.Sweep()
}

func drawObjects(screen *ebiten.Image) {
	// back effects
	for ite := shared.BackEffects.GetIterator(); ite.HasNext(); {
		e := (*effect.Effect)(ite.Next().GetData())
		e.Draw(screen)
	}

	// enemies
	for ite := shared.Enemies.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shooter.Shooter)(obj.GetData())
		e.Draw(screen)
	}

	if player.IsDead() == false {
		player.Draw(screen)
	}

	// player funnels
	for ite := shared.PlayerFunnels.GetIterator(); ite.HasNext(); {
		f := (*shooter.Shooter)(ite.Next().GetData())
		f.Draw(screen)
	}

	// player shots
	for ite := shared.PlayerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		p.Draw(screen)
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
}

func drawMessages(screen *ebiten.Image) {
	if time.Since(dispTextTime).Seconds() <= 1 {
		w := 3 * 24
		h := 24
		shouldPaint := true
		if time.Since(dispTextTime).Seconds() < 0.3 {
			if updateCount%6 > 3 {
				shouldPaint = false
			}
		}
		if shouldPaint {
			paint.DrawText(screen, dispText, screenSize.X/2-w/2,
				h+10, color.White, paint.FontSizeXLarge)
		}
	}
}
