package sprites

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/song9063/ebiutil/geom"
	"github.com/song9063/ebiutil/utils"
)

type Sprite struct {
	geom.Point
	Speed  float64
	target *geom.Point // Destination. stop if nil

	cfg SpriteConfig
}

type SpriteConfig struct {
	DistanceSQArrived float64
}

func DefaultSpriteConfig() SpriteConfig {
	return SpriteConfig{
		DistanceSQArrived: 4.0,
	}
}

func NewSprite(pt geom.Point, speed float64, target *geom.Point, cfg *SpriteConfig) *Sprite {
	config := DefaultSpriteConfig()
	if cfg != nil {
		config = *cfg
	}
	return &Sprite{
		Point:  pt,
		Speed:  speed,
		target: target,
		cfg:    config,
	}
}

func (s *Sprite) SetTarget(p *geom.Point) {
	if p == nil {
		s.target = nil
		return
	}

	// 값만 복사해서 넘거야됨
	copied := *p
	s.target = &copied
}

func (s *Sprite) ClearTarget() {
	s.target = nil
}

func (s *Sprite) IsMoving() bool {
	return s.target != nil
}

func (s *Sprite) Update() bool {
	if s.target == nil {
		return false
	}

	dsq := utils.DistanceSQByPoint(s.Point, *s.target)
	if dsq <= s.cfg.DistanceSQArrived {
		s.X = s.target.X
		s.Y = s.target.Y
		s.target = nil
		return true
	}

	dt := 1.0 / ebiten.ActualTPS()
	dist := math.Sqrt(dsq)
	spd := math.Min(s.Speed*dt, dist)
	s.X += (s.target.X - s.X) / dist * spd
	s.Y += (s.target.Y - s.Y) / dist * spd

	return false
}
