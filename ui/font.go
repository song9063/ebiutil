package ui

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	fontSource *textv2.GoTextFaceSource
	SmallFace  *textv2.GoTextFace
	NormalFace *textv2.GoTextFace
	BigFace    *textv2.GoTextFace
)

func init() {
	s, err := textv2.NewGoTextFaceSource(
		bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontSource = s

	SmallFace = &textv2.GoTextFace{
		Source: fontSource,
		Size:   10,
	}
	NormalFace = &textv2.GoTextFace{
		Source: fontSource,
		Size:   16,
	}
	BigFace = &textv2.GoTextFace{
		Source: fontSource,
		Size:   32,
	}
}

func DrawText(screen *ebiten.Image, str string,
	face *textv2.GoTextFace,
	x, y float64,
	clr color.Color) {
	op := &textv2.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(clr)
	textv2.Draw(screen, str, face, op)
}

func DrawTextCentered(screen *ebiten.Image, str string,
	face *textv2.GoTextFace, cx, cy float64,
	clr color.Color) {
	w, h := textv2.Measure(str, face, 0)
	DrawText(screen, str, face, cx-w/2, cy-h/2, clr)
}
