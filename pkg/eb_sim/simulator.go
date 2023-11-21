package ebsim

import (
	"evs/pkg/config"
	"evs/pkg/population"
	"evs/pkg/ui"
	"fmt"
	"os"

	eb "github.com/hajimehoshi/ebiten/v2"
	ebt "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type EbSim struct {
	Population *population.Population
	// contains all
}

func NewEbSim(n int) *EbSim {
	return &EbSim{
		Population: population.NewPopulation(config.POPULATION_N),
	}
}

func (g *EbSim) Update() error {
	if eb.IsKeyPressed(eb.KeyEscape) {
		os.Exit(0)
	}
	return nil
}

func (g *EbSim) Draw(screen *eb.Image) {
	ebt.DebugPrint(screen, fmt.Sprintf("FPS: %v", eb.ActualFPS()))
	for _, c := range g.Population.Creatures {
		ui.DrawCreature(screen, c)
	}
}

func (g *EbSim) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
