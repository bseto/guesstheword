package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type SimpleGame struct{}

func (s *SimpleGame) Update(screen *ebiten.Image) (err error) {
	err = screen.Fill(color.White)
	if err != nil {
		return
	}
	return
}

func (s *SimpleGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 200
}
