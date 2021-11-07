package shooting

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
	"unsafe"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex"
	"github.com/yohamta/ganim8/v2"

	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/effect"
	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/linkedlist"
	"github.com/yohamta/godanmaku/danmaku/internal/paint"
	"github.com/yohamta/godanmaku/danmaku/internal/quadtree"
	"github.com/yohamta/godanmaku/danmaku/internal/shaders"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/shooter"
	"github.com/yohamta/godanmaku/danmaku/internal/shot"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

const (
	maxPlayerShot   = 300
	maxPlayerFunnel = 20
	maxEnemyShot    = 300
	maxEnemy        = 100
	maxEffects      = 100
	maxBackEffects  = 100
	maxItems        = 10
	quadTreeDepth   = 3
)

type State int

const (
	stateTitle State = iota
	statePlaying
	stateLose
	stateWin
)

type EnemyData struct {
	x, y float64
}

var (
	isInitialized bool
	battleUI      *furex.Flex

	player *shooter.Shooter
	graze  *shooter.Shooter

	state       State
	fld         *field.Field
	enemyQueue  *linkedlist.List
	tmpShooter  *shooter.Shooter
	endTime     time.Time
	killNum     int
	updateCount int

	dispTextTime time.Time
	dispText     string
	dispColor    color.Color

	pShotQuadTree *quadtree.Quadtree
	eShotQuadTree *quadtree.Quadtree

	screenSize   image.Point
	screenCenter image.Point

	// Graze
	grazeCount     int
	grazeBonus     int
	grazeBonusTime int
	earnedBonus    int

	// UI components
	fireButton *FireButton
	joystick   *Joystick
	console    *Console

	offsetImage *ebiten.Image

	// Shader
	shader *ebiten.Shader
)

func startGame() {
	loadResources()
	initAll()
}

type Shooting struct{}

func NewShooting() *Shooting {
	s := &Shooting{}
	return s
}

func (s *Shooting) Layout(width, height int) {
	screenSize = image.Pt(width, height)
	screenCenter.X = screenSize.X / 2
	screenCenter.Y = screenSize.Y/2 - 50
	shared.ScreenSize = screenSize
	isInitialized = true
}

func (s *Shooting) Update() {
	switch state {
	case stateTitle:
		UpdateTitle()
		return
	}

	if isInitialized == false {
		return
	}

	updateCount++

	updateInput()
	updateQuadTree()
	checkCollision()
	updateObjects()

	battleUI.Update()

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
	switch state {
	case stateTitle:
		DrawTitle(screen)
		return
	}
	if isInitialized == false {
		return
	}

	shared.OffsetX = player.GetX() - float64(screenCenter.X)
	shared.OffsetY = player.GetY() - float64(screenCenter.Y)

	drawBackground(offsetImage)
	drawObjects(offsetImage)
	drawMessages(offsetImage)

	switch state {
	case statePlaying:
	case stateLose:
		fallthrough
	case stateWin:
		drawResult(offsetImage)
	}

	if isInitialized {
		battleUI.Draw(offsetImage)
	}

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]interface{}{
		"Time": float32(updateCount),
	}
	op.Images[0] = offsetImage
	screen.DrawRectShader(screenSize.X, screenSize.Y, shader, op)
}

func initObjects() {
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

	// items
	for i := 0; i < maxItems; i++ {
		shared.Items.AddToPool(unsafe.Pointer(shot.NewShot(fld)))
	}

	// player
	player = shooter.NewShooter()

	// graze
	graze = shooter.NewShooter()

	// quad tree
	x0 := fld.GetLeft()
	y0 := fld.GetTop()
	x1 := fld.GetRight()
	y1 := fld.GetBottom()
	pShotQuadTree = quadtree.NewQuadtree(x0, y0, x1, y1, quadTreeDepth)
	eShotQuadTree = quadtree.NewQuadtree(x0, y0, x1, y1, quadTreeDepth)
}

func initAll() {
	initUI()
	initObjects()
	initStage()
	initShaders()
}

func initShaders() {
	var err error
	shader, err = ebiten.NewShader([]byte(shaders.CRT))
	if err != nil {
		panic(err)
	}
}

func initUI() {
	offsetImage = ebiten.NewImage(screenSize.X, screenSize.Y)
	battleUI = furex.NewFlex(screenSize.X, screenSize.Y)

	console = NewConsole()
	battleUI.AddChild(console)

	joystick = NewJoystick()
	battleUI.AddChild(joystick)

	fireButton = NewFireButton()
	battleUI.AddChild(fireButton)
}

func initStage() {
	shared.Enemies.Clean()
	shared.PlayerShots.Clean()
	shared.PlayerFunnels.Clean()
	shared.EnemyShots.Clean()
	shared.Effects.Clean()
	shared.BackEffects.Clean()
	shared.Items.Clean()

	shooter.BuildShooter(shooter.P_ROBO1, player, fld,
		fld.GetCenterX()/2, fld.GetCenterY()/2)
	shooter.BuildGraze(graze, player)

	console.Clear()
	killNum = 0
	enemyQueue.Clear()

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
	screenW := float64(screenSize.X)
	screenH := float64(screenSize.Y)

	spr := sprite.Get("BACKGROUND")
	w, h := spr.Size()
	x, y := screenW/2, screenH/2
	scale := math.Max((screenH / float64(w)), (screenW / float64(h)))

	ganim8.DrawSprite(screen, spr, 0, x, y, 0,
		scale, scale, .5, .5)

	fld.Draw(screen)
}

func drawResult(screen *ebiten.Image) {
	center := screenCenter
	x, y := float64(center.X), float64(center.Y)-50
	spr := sprite.Get("BATTLE_RESULT")
	if state == stateLose {
		ganim8.DrawSprite(screen, spr, 0, x, y, 0, 1, 1, .5, .5)
	} else {
		ganim8.DrawSprite(screen, spr, 1, x, y, 0, 1, 1, .5, .5)
	}
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

	wait := 1
	radius := 100.
	console.Log(Log{
		log:   getText("NEW_ENEMY_APPEAR"),
		color: textColorWarn,
	})
	for i := 0; i < enemyCount; i++ {
		radius += 3
		shooter.BuildShooter(shooter.E_ROBO1, tmpShooter, fld, 0, 0)
		x, y := fld.GetRandamPosition(player.GetX(), player.GetY(), radius)
		x, y = fld.NormalizePosition(x, y, tmpShooter.GetWidth(), tmpShooter.GetHeight())
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
			if enemy.IsDead() {
				continue
			}
			if shot.IsActive() == false {
				continue
			}
			if collision.IsCollideWith(enemy, shot) == false {
				continue
			}

			enemy.AddDamage(1)
			shot.OnHit()
			// sound.PlaySe(sound.SeKindHit2)

			if enemy.IsDead() {
				killNum++
				popItem(enemy.GetX(), enemy.GetY())
				console.Log(Log{
					log:   getText("DESTROY_ENEMY"),
					color: textColorAchievement,
				})
			} else {
				dispTextTime = time.Now()
				dispText = "[命中]"
				dispColor = color.White
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
			if shot.IsGrazed() == false &&
				collision.IsCollideWith(graze, shot) == true {
				checkGraze(shot)
			}

			if collision.IsCollideWith(player, shot) == false {
				continue
			}
			resetGraze()
			player.AddDamage(1)
			shot.OnHit()
			sound.PlaySe(sound.SeKindHit2)

			dispTextTime = time.Now()
			dispText = "[被弾]"
			dispColor = textColorWarn
		}
	}

	// item
	{
		for ite := shared.Items.GetIterator(); ite.HasNext(); {
			i := (*shot.Shot)(ite.Next().GetData())
			if player.IsDead() {
				break
			}
			if i.IsActive() == false {
				continue
			}
			if collision.IsCollideWith(player, i) == false {
				continue
			}
			i.SetInactive()
			getItem(i.GetItemKind())
			sound.PlaySe(sound.SeKindItemGet)
		}
	}
}

func resetGraze() {
	grazeCount = 0
	grazeBonusTime = 0
}

func checkGraze(s *shot.Shot) {
	if shared.Enemies.GetActiveNum() < 4 &&
		shared.EnemyShots.GetActiveNum() < 32 {
		return
	}
	grazeCount++
	grazeBonusTime = 60
	grazeBonus = grazeBonus + 1
	if grazeBonus > 100 {
		grazeBonus = 100
	}
	earnedBonus += grazeBonus
	s.SetGrazed()

	x := (s.GetX() + graze.GetX()) / 2
	y := (s.GetY() + graze.GetY()) / 2
	effect.CreateGraze(x, y)
}

func getItem(itemKind shot.ItemKind) {
	switch itemKind {
	case shot.ItemKindPowerUp:
		addFunnel()
	case shot.ItemKindRecovery:
		player.Recovery()
	}
}

func loadResources() {
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
			if player.Fire() {
				sound.PlaySe(sound.SeKindShot)
			}
		}
	}

	// graze
	graze.Update()

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

	// items
	for ite := shared.Items.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		i := (*shot.Shot)(obj.GetData())
		if i.IsActive() == false {
			obj.SetInactive()
			continue
		}
		i.Update()
	}

	shared.EnemyShots.Sweep()
	shared.PlayerShots.Sweep()
	shared.PlayerFunnels.Sweep()
	shared.Enemies.Sweep()
	shared.Effects.Sweep()
	shared.BackEffects.Sweep()
	shared.Items.Sweep()
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

	// items
	for ite := shared.Items.GetIterator(); ite.HasNext(); {
		i := (*shot.Shot)(ite.Next().GetData())
		i.Draw(screen)
	}
}

func popItem(x, y float64) {
	a := shared.Enemies.GetActiveNum()
	if a > 20 {
		a = 20
	}
	chance := a + grazeCount + 10
	if chance >= 50 {
		chance = 50
	}
	if rand.Intn(100) > chance {
		return
	}
	console.Log(Log{
		log:   getText("ITEM_APPEAR"),
		color: textColorInfo,
	})
	if rand.Float64() > 0.4 {
		shot.RecoveryItem(x, y, rand.Intn(359))
	} else {
		shot.PowerUpItem(x, y, rand.Intn(359))
	}
}

func addFunnel() {
	funnel := (*shooter.Shooter)(shared.PlayerFunnels.CreateFromPool())
	if funnel == nil {
		return
	}
	shooter.BuildFunnel(funnel, player, fld)
}

func drawMessages(screen *ebiten.Image) {
	if time.Since(dispTextTime).Seconds() <= 1 {
		w := 3 * 24
		shouldPaint := true
		if time.Since(dispTextTime).Seconds() < 0.3 {
			if updateCount%6 > 3 {
				shouldPaint = false
			}
		}
		if shouldPaint {
			paint.DrawText(screen, dispText, screenSize.X/4*3-w/2,
				screenSize.Y-130, dispColor, paint.FontSizeXLarge)
		}
	}
}
