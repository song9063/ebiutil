package effects

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/song9063/ebiutil/geom"
	"github.com/song9063/ebiutil/utils"
)

type EffectType int

const (
	EffectScale EffectType = iota // Scale
	EffectShake                   // Shake
	EffectSlide                   // Slide
)

type BannerContent interface {
	Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions)
	Size() (width, height float64)
}

type TextContent struct {
	Text  string
	Font  *text.GoTextFace
	Color color.Color

	cached *ebiten.Image
}

type ImageContent struct {
	Image *ebiten.Image
}

type BannerConfig struct {
	Effect   EffectType
	Duration float64
	FadeTime float64 // Fade out

	// EffectScale
	ScaleStart float64 // Default 0.3

	// EffectShake
	ShakeSpeed     float64 // Default 40
	ShakeAmplitude float64 // Default 8

	// EffectSlide
	SlideStartX float64 // default -20
	SlideSpeed  float64 // default 0.2
}

func DefaultBannerConfig() BannerConfig {
	return BannerConfig{
		Effect:         EffectScale,
		Duration:       2.0,
		FadeTime:       0.5,
		ScaleStart:     0.3,
		ShakeSpeed:     40,
		ShakeAmplitude: 8,
		SlideStartX:    -200,
		SlideSpeed:     0.2,
	}
}

type Banner struct {
	content BannerContent
	cfg     BannerConfig
	active  bool
	timer   float64

	scale  float64
	offset geom.Point
	alpha  float32
}

func (b *Banner) Show(content BannerContent, cfg BannerConfig) {
	b.content = content
	b.cfg = cfg
	b.active = true
	b.timer = cfg.Duration
	b.alpha = 1.0
	switch cfg.Effect {
	case EffectScale:
		b.scale = b.cfg.ScaleStart
	case EffectShake:
		b.scale = 1.0
	case EffectSlide:
		b.scale = 1.0
		b.offset.X = b.cfg.SlideStartX
	}
}

func (b *Banner) Update() {
	if !b.active {
		return
	}

	dt := utils.ActualDeltaTime()
	b.timer -= dt

	switch b.cfg.Effect {
	case EffectScale:
		// Lerp: 현재값 = 현재값 + (목표값 - 현재값) * 속도
		b.scale += (1.0 - b.scale) * 0.2
	case EffectShake:
		// sin으로 좌우로 흔들기
		amplitude := b.cfg.ShakeAmplitude * float64(b.alpha)
		b.offset.X = math.Sin(b.timer*b.cfg.ShakeSpeed) * amplitude
	case EffectSlide:
		b.offset.X += (0 - b.offset.X) * b.cfg.SlideSpeed
	}

	// Fadeout
	if b.timer < b.cfg.FadeTime {
		b.alpha = float32(b.timer / b.cfg.FadeTime)
	}

	// Finish
	if b.timer <= 0 {
		b.active = false
	}
}

func (b *Banner) Draw(screen *ebiten.Image, at geom.Point) {
	if !b.active {
		return
	}

	w, h := b.content.Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-w/2, -h/2)
	op.GeoM.Scale(b.scale, b.scale)
	op.GeoM.Translate(at.X+b.offset.X,
		at.Y+b.offset.Y)
	op.ColorScale.ScaleAlpha(b.alpha)
	b.content.Draw(screen, op)
}

func (b *Banner) IsActive() bool {
	return b.active
}

func (t *TextContent) Size() (float64, float64) {
	w, h := text.Measure(t.Text, t.Font, t.Font.Size)
	return w, h
}

func (t *TextContent) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	if t.cached == nil {
		w, h := t.Size()
		t.cached = ebiten.NewImage(int(w), int(h))
		textOp := &text.DrawOptions{}
		if t.Color != nil {
			textOp.ColorScale.ScaleWithColor(t.Color)
		}
		text.Draw(t.cached, t.Text, t.Font, textOp)
	}

	screen.DrawImage(t.cached, op)
}

func (i *ImageContent) Size() (float64, float64) {
	b := i.Image.Bounds()
	return float64(b.Dx()), float64(b.Dy())
}

func (i *ImageContent) Draw(screen *ebiten.Image,
	op *ebiten.DrawImageOptions) {
	screen.DrawImage(i.Image, op)
}
