package creatures

import (
	"image/color"

	"golang.org/x/image/colornames"
)

type Creature struct {
	X   int64
	Y   int64
	Age int64
}

func NewCreature(x, y int64) *Creature {
	return &Creature{
		X:   x,
		Y:   y,
		Age: 0,
	}
}

func (c *Creature) Color() color.Color {
	// TODO: make dependent on some creature param
	return colornames.Green
}
