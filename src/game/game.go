package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameInterface interface {
}

type Game struct {
	screenHeight int
	screenWidth  int
}

func (g *Game) getScreenHeight() int {
	return g.screenHeight
}

func (g *Game) getScreenWidth() int {
	return g.screenWidth
}

func (g *Game) SetScreenHeight(height int) error {
	g.screenHeight = height
	return nil
}

func (g *Game) SetScreenWidth(width int) error {
	g.screenWidth = width
	return nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	squareColor := color.RGBA{255, 0, 0, 255}

	// Calculando a posição para centralizar o quadrado na tela.
	x := (g.screenWidth - 50) / 2
	y := (g.screenHeight - 50) / 2

	// Desenhando o quadrado na tela.
	vector.DrawFilledRect(screen, float32(x), float32(y), 50, 50, squareColor, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
