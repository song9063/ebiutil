package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/song9063/ebiutil/geom"
	"image/color"
)

var (
	DefaultButtonBGColor   = color.RGBA{0x0, 0x50, 0x90, 0xff}
	DefaultButtonTextColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

type Button struct {
	geom.Rect
	Text string

	BackgroundColor color.Color
	TextColor       color.Color
}

func NewButton(rect geom.Rect, text string) *Button {
	return &Button{
		Rect:            rect,
		Text:            text,
		BackgroundColor: DefaultButtonBGColor,
		TextColor:       DefaultButtonTextColor,
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	vector.FillRect(screen,
		float32(b.X), float32(b.Y),
		float32(b.W), float32(b.H),
		b.BackgroundColor, false)

	DrawTextCentered(screen, b.Text, NormalFace,
		b.CenterX(), b.CenterY(), b.TextColor)
}

func (b *Button) IsPressed(bt ebiten.MouseButton) bool {
	mx, my := ebiten.CursorPosition()
	p := geom.NewPointFromInt(mx, my)
	return b.HitTest(p) &&
		ebiten.IsMouseButtonPressed(bt)
}

func (b *Button) IsJustPressed(bt ebiten.MouseButton) bool {
	mx, my := ebiten.CursorPosition()
	p := geom.NewPointFromInt(mx, my)
	return b.HitTest(p) &&
		inpututil.IsMouseButtonJustPressed(bt)
}
