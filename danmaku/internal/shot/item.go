package shot

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
	"github.com/yohamta/godanmaku/danmaku/internal/collision"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type ItemKind int

const (
	ItemKindPowerUp = iota
	ItemKindRecovery
)

type item struct {
	baseController
}

func (c *item) init(s *Shot) {
	switch s.itemKind {
	case ItemKindPowerUp:
		s.spr = sprite.Get("ITEM_P")
		s.collisionBox = collision.GetCollisionBox("ITEM_P")
	case ItemKindRecovery:
		s.spr = sprite.Get("ITEM_LIFE")
		s.collisionBox = collision.GetCollisionBox("ITEM_LIFE")
	}
	s.sprIndex = 0
	s.setSize(float64(s.spr.W()), float64(s.spr.H()))
	s.setSpeed(ShotSpd1, s.degree)
}

func (c *item) draw(s *Shot, screen *ebiten.Image) {
	x, y := s.x-shared.OffsetX, s.y-shared.OffsetY
	ganim8.DrawSprite(screen, s.spr, 0, x, y, 0, 1, 1, .5, .5)
}
