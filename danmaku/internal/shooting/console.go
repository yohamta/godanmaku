package shooting

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miyahoyo/godanmaku/danmaku/internal/paint"
)

const maxLogs = 5

type Log struct {
	log   string
	color color.Color
}

type Console struct {
	logs  [maxLogs]Log
	first int
	size  int
}

func NewConsole() *Console {
	c := &Console{}

	return c
}

func (c *Console) Log(log Log) {
	if c.size == maxLogs {
		if c.first == maxLogs-1 {
			c.first = 0
		} else {
			c.first++
		}
	} else {
		c.size++
	}
	c.logs[(c.first+c.size-1)%(maxLogs)] = log
}

func (c *Console) Size() (int, int) {
	return screenSize.X, screenSize.Y
}

func (c *Console) Position() (int, int) {
	return 0, 0
}

func (c *Console) Clear() {
	c.first = 0
	c.size = 0
}

func (c *Console) Draw(screen *ebiten.Image, frame image.Rectangle) {
	for i := 0; i < c.size; i++ {
		log := c.logs[(c.first+i)%(maxLogs)]
		paint.DrawTextWithOptions(screen, log.log, frame.Min.X+4,
			frame.Min.Y+i*12+4, paint.DrawTextOptions{
				Color:    log.color,
				Width:    screenSize.X,
				HAligh:   paint.HAlignLeft,
				FontSize: paint.FontSizeSmall,
			})
	}
}
