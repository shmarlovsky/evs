package main

import (
	"evs/pkg/config"
	ebsim "evs/pkg/eb_sim"
	"log"
	"math/rand"
	"time"

	eb "github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_X = config.MAX_X
	SCREEN_Y = config.MAX_Y
)

func main() {
	rand.Seed(time.Now().UnixNano())

	eb.SetWindowSize(SCREEN_X, SCREEN_Y)
	eb.SetWindowTitle("Watch them evolve")
	s := ebsim.NewEbSim(10)
	if err := eb.RunGame(s); err != nil {
		log.Fatal(err)
	}
}
