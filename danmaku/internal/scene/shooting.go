package scene

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/yohamta/godanmaku/danmaku/internal/field"
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

// PlayerShooter represents interface of Player Weapon
type PlayerShooter interface {
	Shot(x, y float64, degree int, playerShots []*shot.Shot)
}

var (
	screenWidth  = 0
	screenHeight = 0

	input        *inputs.Input
	currentField *field.Field

	uiBackground      *ui.Box
	uiBackgroundColor = color.RGBA{0x00, 0x00, 0x00, 0xff}

	player       *shooter.Player
	playerWeapon PlayerShooter

	playerShots [maxPlayerShot]*shot.Shot
	enemyShots  [maxEnemyShot]*shot.Shot
	enemies     [maxEnemy]*shooter.Enemy
	hitEffects  [maxHitEffects]*effects.Hit
	explosions  [maxExplosions]*effects.Explosion

	state gameState = gameStateLoading
)

// Shooting represents shooting scene
type Shooting struct {
}

// NewShootingOptions represents options for New func
type NewShootingOptions struct {
	ScreenWidth  int
	ScreenHeight int
}

// NewShooting returns new Shooting struct
func NewShooting(options NewShootingOptions) *Shooting {
	stg := &Shooting{}

	screenWidth = options.ScreenWidth
	screenHeight = options.ScreenHeight

	state = gameStateLoading
	initGame()
	state = gameStatePlaying

	return stg
}

func initGame() {
	rand.Seed(time.Now().Unix())
	input = inputs.NewInput(screenWidth, screenHeight)
	currentField = field.NewField()
	uiBackground = ui.NewBox(0, int(currentField.GetBottom()),
		screenWidth, screenHeight-int(currentField.GetBottom()-currentField.GetTop()),
		uiBackgroundColor)

	// player
	player = shooter.NewPlayer()
	player.SetMainWeapon(weapon.NewNormal(shot.KindPlayerNormal))
	player.SetField(currentField)

	// enemies
	for i := 0; i < len(enemies); i++ {
		enemies[i] = shooter.NewEnemy()
		enemies[i].SetField(currentField)
	}

	// shots
	for i := 0; i < len(playerShots); i++ {
		playerShots[i] = shot.NewShot()
		playerShots[i].SetField(currentField)
	}

	// enemyShots
	for i := 0; i < len(enemyShots); i++ {
		enemyShots[i] = shot.NewShot()
		enemyShots[i].SetField(currentField)
	}

	// effects
	for i := 0; i < len(hitEffects); i++ {
		hitEffects[i] = effects.NewHit()
	}
	for i := 0; i < len(explosions); i++ {
		explosions[i] = effects.NewExplosion()
	}

	// Setup stage
	initEnemies()
}

// Update updates the scene
func (stg *Shooting) Update() {
	input.Update()

	checkCollision()

	// player
	if player.IsDead() == false {
		player.Move(input.Horizontal, input.Vertical, input.Fire)
		if input.Fire {
			player.FireWeapon(playerShots[:])
		}
	}

	// player shots
	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Move()
	}

	// enemy shots
	for i := 0; i < len(enemyShots); i++ {
		e := enemyShots[i]
		if e.IsActive() == false {
			continue
		}
		e.Move()
	}

	// enemies
	for i := 0; i < len(enemies); i++ {
		e := enemies[i]
		if e.IsActive() == false {
			continue
		}
		e.Move()
		if player.IsDead() == false {
			e.FireWeapon(enemyShots[:])
		}
	}

	// hitEffects
	for i := 0; i < len(hitEffects); i++ {
		h := hitEffects[i]
		if h.IsActive() == false {
			continue
		}
		h.Update()
	}

	// explosions
	for i := 0; i < len(explosions); i++ {
		e := explosions[i]
		if e.IsActive() == false {
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
	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Draw(screen)
	}

	// enemies
	for i := 0; i < len(enemies); i++ {
		e := enemies[i]
		if e.IsActive() == false {
			continue
		}
		e.Draw(screen)
	}

	if player.IsDead() == false {
		player.Draw(screen)
	}

	// enemy shots
	for i := 0; i < len(enemyShots); i++ {
		e := enemyShots[i]
		if e.IsActive() == false {
			continue
		}
		e.Draw(screen)
	}

	// explosions
	for i := 0; i < len(explosions); i++ {
		e := explosions[i]
		if e.IsActive() == false {
			continue
		}
		e.Draw(screen)
	}

	// hitEffects
	for i := 0; i < len(hitEffects); i++ {
		h := hitEffects[i]
		if h.IsActive() == false {
			continue
		}
		h.Draw(screen)
	}

	uiBackground.Draw(screen)
	input.Draw(screen)
}

func initEnemies() {
	enemyCount := 20

	for i := 0; i < enemyCount; i++ {
		enemy := enemies[i]
		enemy.Init()
		enemy.SetTarget(player)
	}
}

func checkCollision() {
	// player shots
	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		for j := 0; j < len(enemies); j++ {
			e := enemies[j]
			if e.IsActive() == false {
				continue
			}
			if e.IsCollideWith(p.GetEntity()) == false {
				continue
			}
			e.AddDamage(1)
			p.SetActive(false)
			createHitEffect(p.GetX(), p.GetY())
			if e.IsDead() {
				createExplosion(e.GetX(), e.GetY())
			}
		}
	}

	// enemy shots
	if player.IsDead() == false {
		for i := 0; i < len(enemyShots); i++ {
			e := enemyShots[i]
			if e.IsActive() == false {
				continue
			}
			if player.IsCollideWith(e.GetEntity()) == false {
				continue
			}
			player.AddDamage(1)
			e.SetActive(false)
			createHitEffect(player.GetX(), player.GetY())
			if player.IsDead() {
				createExplosion(player.GetX(), player.GetY())
			}
		}
	}
}

func createHitEffect(x, y float64) {
	for i := 0; i < len(hitEffects); i++ {
		h := hitEffects[i]
		if h.IsActive() {
			continue
		}
		h.StartEffect(x, y)
		break
	}
}

func createExplosion(x, y float64) {
	for i := 0; i < len(explosions); i++ {
		e := explosions[i]
		if e.IsActive() {
			continue
		}
		e.StartEffect(x, y)
		break
	}
}
