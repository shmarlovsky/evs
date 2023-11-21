package ui

import (
	"evs/pkg/creatures"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	CREATURE_X = 3
	CREATURE_Y = 3
)

func DrawCreature(dst *eb.Image, creature *creatures.Creature) {
	img := eb.NewImage(CREATURE_X, CREATURE_Y)
	img.Fill(creature.Color())
	geoM := eb.GeoM{}
	geoM.Translate(float64(creature.X), float64(creature.Y))
	dst.DrawImage(img, &eb.DrawImageOptions{
		GeoM: geoM,
	})
}
