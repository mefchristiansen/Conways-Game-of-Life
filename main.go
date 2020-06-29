package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/mefchristiansen/conways-game-of-life/gameoflife"
)

func main() {
	game, err := gameoflife.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetMaxTPS(20)
	ebiten.SetWindowSize(gameoflife.ScreenDimension, gameoflife.ScreenDimension)
	ebiten.SetWindowTitle("Conway's Game of Life")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
