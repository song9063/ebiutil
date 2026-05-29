package ui

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/song9063/ebiutil/geom"
)

type TextWithGeom struct {
	Rect geom.Rect
	text string
	Face *textv2.GoTextFace
}

func MakeTextWithGeom(txt string, face *textv2.GoTextFace,
	pos geom.Point) *TextWithGeom {
	w, h := textv2.Measure(txt, face, 0)
	return &TextWithGeom{
		Rect: geom.Rect{
			Point: pos,
			W:     w, H: h,
		},
		text: txt,
		Face: face,
	}
}
func MakeTextWithGeomCenter(txt string, face *textv2.GoTextFace,
	center geom.Point) *TextWithGeom {
	txtW := MakeTextWithGeom(txt, face, center)
	// txtW.Rect.SetCenter(center)
	txtW.SetCenter(center)
	return txtW
}
func (t *TextWithGeom) SetCenter(center geom.Point) {
	t.Rect.SetCenter(center)
}
func (t *TextWithGeom) SetText(txt string) {
	t.text = txt
	w, h := textv2.Measure(txt, t.Face, 0)
	center := t.Rect.Center()
	t.Rect.W = w
	t.Rect.H = h
	t.Rect.SetCenter(center)
}
func (t *TextWithGeom) GetText() string {
	return t.text
}

var (
	fontSource    *textv2.GoTextFaceSource
	SmallFace     *textv2.GoTextFace
	LessSmallFace *textv2.GoTextFace
	NormalFace    *textv2.GoTextFace
	BigFace       *textv2.GoTextFace
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
	LessSmallFace = &textv2.GoTextFace{
		Source: fontSource,
		Size:   13,
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
