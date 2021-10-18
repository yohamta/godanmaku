package shot

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miyahoyo/godanmaku/danmaku/internal/collision"
	"github.com/miyahoyo/godanmaku/danmaku/internal/shared"
	"github.com/miyahoyo/godanmaku/danmaku/internal/sprite"
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
		s.spr = sprite.ItemP
		s.collisionBox = collision.GetCollisionBox("ITEM_P")
	case ItemKindRecovery:
		s.spr = sprite.ItemL
		s.collisionBox = collision.GetCollisionBox("ITEM_LIFE")
	}
	s.sprIndex = 0
	s.setSize(float64(s.spr.GetWidth()), float64(s.spr.GetHeight()))
	s.setSpeed(ShotSpd1, s.degree)
}

func (c *item) draw(s *Shot, screen *ebiten.Image) {
	spr := s.spr
	spr.SetPosition(s.x-shared.OffsetX, s.y-shared.OffsetY)
	spr.SetIndex(s.sprIndex)
	spr.Draw(screen)
}
