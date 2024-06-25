package main

import (
	"go_plataformer/src/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game_instance := game.Game{
		screenHeight: 600,
		screenWidth:  600,
	}
	ebiten.SetWindowSize(game_instance.screenWidth, game_instance.screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game_instance); err != nil {
		log.Fatal(err)
	}
}
