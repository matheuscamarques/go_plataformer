package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/matheuscamarques/go_plataformer/src/game"
)

func main() {
	game_instance := &game.Game{}
	game_instance.SetScreenHeight(600)
	game_instance.SetScreenWidth(600)

	ebiten.SetWindowSize(game_instance.getScreenWidth(), game_instance.getScreenHeight())
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game_instance); err != nil {
		log.Fatal(err)
	}
}
