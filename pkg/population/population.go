package population

import (
	"evs/pkg/config"
	"evs/pkg/creatures"
	"math/rand"
)

type Population struct {
	Creatures []*creatures.Creature
}

func NewPopulation(n int) *Population {
	c := make([]*creatures.Creature, n)
	for i := 0; i < n; i++ {
		c[i] = creatures.NewCreature(
			RandomFromInterval(0, config.MAX_X),
			RandomFromInterval(0, config.MAX_Y),
		)
	}
	return &Population{
		Creatures: c,
	}
}

func RandomFromInterval(min, max int64) int64 {
	r := rand.Int63n(max)
	if r < min {
		r += min
	}
	return r
}
