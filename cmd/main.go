package main

import (
	"log"

	"github.com/bseto/guesstheword/pkg/game"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowSize(400, 200)
	ebiten.SetWindowTitle("Hello Sarah")
	err := ebiten.RunGame(&game.SimpleGame{})
	if err != nil {
		log.Fatalf("unable to run game: %v", err)
	}
}
