package shooting

import (
	"image/color"
	"math/rand"
	"time"
	"unsafe"

	"github.com/yohamta/godanmaku/danmaku/internal/flyweight"

	"github.com/yohamta/godanmaku/danmaku/internal/field"
	"github.com/yohamta/godanmaku/danmaku/internal/util"
	"github.com/yohamta/godanmaku/danmaku/internal/weapon"

	"github.com/yohamta/godanmaku/danmaku/internal/effects"
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
	maxHitEffects = 30
	maxExplosions = 30
)

type gameState int

const (
	gameStateLoading gameState = iota
	gameStatePlaying
)

var (
	input        *inputs.Input
	currentField *field.Field

	background      *ui.Box
	backgroundColor = color.RGBA{0x00, 0x00, 0x00, 0xff}

	player *shooter.Player

	playerShots *flyweight.Factory
	enemyShots  *flyweight.Factory
	enemies     *flyweight.Factory
	hitEffects  *flyweight.Factory
	explosions  *flyweight.Factory

	state gameState = gameStateLoading
)

// Shooting represents shooting scene
type Shooting struct {
	screenWidth  int
	screenHeight int
}

// NewShooting returns new Shooting struct
func NewShooting(screenWidth, screenHeight int) *Shooting {
	stg := &Shooting{}

	stg.screenWidth = screenWidth
	stg.screenHeight = screenHeight

	stg.initGame()

	return stg
}

func (stg *Shooting) initGame() {
	state = gameStateLoading
	rand.Seed(time.Now().Unix())
	input = inputs.NewInput(stg.screenWidth, stg.screenHeight)
	currentField = field.NewField()

	background = ui.NewBox(0, int(currentField.GetBottom()),
		stg.screenWidth,
		stg.screenHeight-int(currentField.GetBottom()-currentField.GetTop()),
		backgroundColor)

	// player
	player = shooter.NewPlayer(currentField)
	player.Init()
	player.SetMainWeapon(weapon.NewNormal(shot.KindPlayerNormal))
	player.SetField(currentField)

	// enemies
	enemies = flyweight.NewFactory()
	for i := 0; i < maxEnemy; i++ {
		enemies.AddToPool(unsafe.Pointer(shooter.NewEnemy(currentField)))
	}

	// shots
	playerShots = flyweight.NewFactory()
	for i := 0; i < maxPlayerShot; i++ {
		playerShots.AddToPool(unsafe.Pointer(shot.NewShot(currentField)))
	}

	// enemyShots
	enemyShots = flyweight.NewFactory()
	for i := 0; i < maxEnemyShot; i++ {
		enemyShots.AddToPool(unsafe.Pointer(shot.NewShot(currentField)))
	}

	// effects
	hitEffects = flyweight.NewFactory()
	for i := 0; i < maxHitEffects; i++ {
		hitEffects.AddToPool(unsafe.Pointer(effects.NewHit()))
	}
	explosions = flyweight.NewFactory()
	for i := 0; i < maxExplosions; i++ {
		explosions.AddToPool(unsafe.Pointer(effects.NewExplosion()))
	}

	// Setup stage
	initEnemies()
	state = gameStatePlaying
}

// Update updates the scene
func (stg *Shooting) Update() {
	input.Update()

	checkCollision()

	// player
	if player.IsDead() == false {
		player.Move(input.Horizontal, input.Vertical, input.Fire)
		if input.Fire {
			player.FireWeapon(playerShots)
		}
	}

	// player shots
	for ite := playerShots.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		p := (*shot.Shot)(obj.GetData())
		if p.IsActive() == false {
			obj.SetInactive()
			continue
		}
		p.Move()
	}
	playerShots.Sweep()

	// enemy shots
	for ite := enemyShots.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shot.Shot)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			continue
		}
		e.Move()
	}
	enemyShots.Sweep()

	// enemies
	for ite := enemies.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shooter.Enemy)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			continue
		}
		e.Move()
		if player.IsDead() == false {
			e.FireWeapon(enemyShots)
		}
	}

	// hitEffects
	for ite := hitEffects.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		h := (*effects.Hit)(obj.GetData())
		if h.IsActive() == false {
			obj.SetInactive()
			continue
		}
		h.Update()
	}

	// explosions
	for ite := explosions.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*effects.Explosion)(obj.GetData())
		if e.IsActive() == false {
			obj.SetInactive()
			continue
		}
		e.Update()
	}
}

// Draw draws the scene
func (stg *Shooting) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})

	currentField.Draw(screen)

	// player shots
	for ite := playerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		if p.IsActive() == false {
			continue
		}
		p.Draw(screen)
	}

	// enemies
	for ite := enemies.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		e := (*shooter.Enemy)(obj.GetData())
		if e.IsActive() == false {
			continue
		}
		e.Draw(screen)
	}

	if player.IsDead() == false {
		player.Draw(screen)
	}

	// enemy shots
	for ite := enemyShots.GetIterator(); ite.HasNext(); {
		e := (*shot.Shot)(ite.Next().GetData())
		if e.IsActive() == false {
			continue
		}
		e.Draw(screen)
	}

	// explosions
	for ite := explosions.GetIterator(); ite.HasNext(); {
		e := (*effects.Explosion)(ite.Next().GetData())
		if e.IsActive() == false {
			continue
		}
		e.Draw(screen)
	}

	// hitEffects
	for ite := hitEffects.GetIterator(); ite.HasNext(); {
		h := (*effects.Hit)(ite.Next().GetData())
		if h.IsActive() == false {
			continue
		}
		h.Draw(screen)
	}

	background.Draw(screen)
	input.Draw(screen)
}

func initEnemies() {
	enemyCount := 20

	for i := 0; i < enemyCount; i++ {
		enemy := (*shooter.Enemy)(enemies.CreateFromPool())
		if enemy == nil {
			return
		}
		enemy.Init()
		enemy.SetTarget(player)
	}
}

func checkCollision() {
	// player shots
	for ite := playerShots.GetIterator(); ite.HasNext(); {
		p := (*shot.Shot)(ite.Next().GetData())
		if p.IsActive() == false {
			continue
		}
		for ite2 := enemies.GetIterator(); ite2.HasNext(); {
			e := (*shooter.Enemy)(ite2.Next().GetData())
			if e.IsActive() == false {
				continue
			}
			if util.IsCollideWith(e, p) == false {
				continue
			}
			e.AddDamage(1)
			p.OnHit()
			createHitEffect(p.GetX(), p.GetY())
			if e.IsDead() {
				createExplosion(e.GetX(), e.GetY())
			}
		}
	}

	// enemy shots
	if player.IsDead() == false {
		for ite := enemyShots.GetIterator(); ite.HasNext(); {
			e := (*shot.Shot)(ite.Next().GetData())
			if e.IsActive() == false {
				continue
			}
			if util.IsCollideWith(player, e) == false {
				continue
			}
			player.AddDamage(1)
			e.OnHit()
			createHitEffect(player.GetX(), player.GetY())
			if player.IsDead() {
				createExplosion(player.GetX(), player.GetY())
			}
		}
	}
}

func createHitEffect(x, y float64) {
	h := (*effects.Hit)(hitEffects.CreateFromPool())
	if h == nil {
		return
	}
	h.StartEffect(x, y)
}

func createExplosion(x, y float64) {
	e := (*effects.Explosion)(explosions.CreateFromPool())
	if e == nil {
		return
	}
	e.StartEffect(x, y)
}
