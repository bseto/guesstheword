package game

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/bseto/guesstheword/pkg/resources/fonts"
	"github.com/hajimehoshi/ebiten"
)

// Too lazy to use uber.Fx here
var normalFont font.Face

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular)
	if err != nil {
		log.Fatalf("unable to Parse: %v", err)
	}
	normalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatalf("unable to setup NewFace: %v", err)
	}

}

type Letter interface {
	Draw(canvas *ebiten.Image)
}

type LetterImpl struct {
	normalFont font.Face
}

func NewLetter() Letter {
	return &LetterImpl{
		normalFont: normalFont,
	}
}

func (l *LetterImpl) Draw(canvas *ebiten.Image) {
}
