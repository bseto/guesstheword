package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowSize(400, 200)
	ebiten.SetWindowTitle("Hello Sarah")
	err := ebiten.RunGame(&SimpleGame{})
	if err != nil {
		log.Fatalf("unable to run game: %v", err)
	}
}

type SimpleGame struct{}

func (s *SimpleGame) Update(screen *ebiten.Image) error {
	return nil
}

func (s *SimpleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 200
}
