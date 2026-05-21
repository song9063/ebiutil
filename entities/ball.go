package entities

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/song9063/ebiutil/geom"
)

// Usage
// func (b *BattleRunScene) startBallFly(hitType common.PitchResult_t) {
// 	angleRad := utils.RandomRadianInRange(45, 90)
// 	if hitType == common.PitchResult_Foul {
// 		angleRad = utils.RandomRadianInRange(135, 270)
// 	}
// 	ballParams := common.BallFlyParamsByResult[hitType]

// 	vz := utils.RandomFloat64(ballParams.VZMin,
// 		ballParams.VZMax)
// 	horSpeed := utils.RandomFloat64(ballParams.HorSpeedMin,
// 		ballParams.HorSpeedMax)
// 	vx, vy := utils.VelocityFromAngle(angleRad, horSpeed)
// 	ballCfg := entities.DefaultBallConfig()

// 	b.fieldState.Ball = entities.NewBall(
// 		b.fieldState.Bases[0].ToPoint64(),
// 		common.BALL_RADIUS,
// 		vx, vy, vz,
// 		common.GRAVITY, ballCfg,
// 	)
// }

type Ball struct {
	geom.Point         // Position(X,Y)
	Z          float64 // Height(0=Ground)

	Radius float64

	VX, VY   float64 // Horizontal speed
	VZ       float64 // Vertical speed
	Gravity  float64
	Bouncing bool

	// Parameters
	cfg BallConfig
}

type BallConfig struct {
	// Physics
	BounceDecay   float64 // 바운스 감쇠량
	FrictionDecay float64 // 마찰 감쇠량
	StopThreshold float64 // 멈춤 판정 속도

	//
	Color color.Color
}

func DefaultBallConfig() BallConfig {
	return BallConfig{
		BounceDecay:   0.5,
		FrictionDecay: 0.6,
		StopThreshold: 0.1,
		Color:         color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
}

func NewBall(pos geom.Point, radius, vx, vy, vz, gravity float64, cfg BallConfig) *Ball {
	return &Ball{
		Point:    pos,
		Z:        0,
		Radius:   radius,
		VX:       vx,
		VY:       vy,
		VZ:       vz,
		Gravity:  gravity,
		Bouncing: false,
		cfg:      cfg,
	}
}

func (b *Ball) Stop() {
	b.VX = 0
	b.VY = 0
	b.VZ = 0
	b.Z = 0
	b.Bouncing = false
}

// return true if stopped at target
func (b *Ball) Update() bool {
	dt := 1.0 / ebiten.ActualTPS()

	// X,Y,Z 변화
	b.X += b.VX * dt
	b.Y += b.VY * dt
	b.Z += b.VZ * dt

	// 중력에 의해 수직속도 점점 감소, 음수가 되면 아래로 떨어지게됨
	b.VZ -= b.Gravity * dt

	// 지면에 떨어짐
	if b.Z <= 0 {
		b.Z = 0 // 땅 밑으로 더 내려가지않게

		// 수직 속도를 반대로 감쇠시킴
		// 곱하는 값은 튀어오르는 힘(1이면 영원히 튀게됨)
		b.VZ = -b.VZ * b.cfg.BounceDecay

		// 수평속도 떨어트리면서 굴러감
		b.VX *= b.cfg.FrictionDecay
		b.VY *= b.cfg.FrictionDecay

		// 속도가 많이 떨어지면 멈춤
		if math.Abs(b.VZ) < b.cfg.StopThreshold {
			b.VZ = 0
			b.Bouncing = false // 튀는 힘 없어짐
		} else {
			b.Bouncing = true // 통통뛰는중
		}

		// 지면에서 굴러가는 동안 속도 감쇠
		if b.VZ == 0 {
			b.VX *= b.cfg.FrictionDecay
			b.VY *= b.cfg.FrictionDecay
		}
	}

	// 완전히 멈췄는지 판정
	if math.Abs(b.VX) < b.cfg.StopThreshold && math.Abs(b.VY) < b.cfg.StopThreshold && b.Z == 0 {
		return true
	}

	return false
}

func (b *Ball) Draw(screen *ebiten.Image) {
	radius := float32(b.Radius + b.Z*0.05)
	vector.FillCircle(screen,
		float32(b.X), float32(b.Y),
		radius,
		b.cfg.Color, false)
}

// 볼 착지지점 계산
// [착지 시간 계산]
// Z = VZ * t - 0.5 * Gravity * t²

// 착지 = Z가 0이 될 때
// 0 = VZ * t - 0.5 * Gravity * t²
// 0 = t * (VZ - 0.5 * Gravity * t)
// t = 0 또는 t = (2 * VZ) / Gravity

// → 착지 시간 = 2 * VZ / Gravity

// [착지 위치 계산]
// 착지X = 출발X + VX * 착지시간
// 착지Y = 출발Y + VY * 착지시간
func (b *Ball) LandingPosition() (geom.Point, bool) {
	if b.Gravity <= 0 {
		return geom.Point{}, false
	}

	if b.VZ <= 0 {
		return b.Point, true // 현재 위치가 착지 지점
	}

	// 착지까지 걸리는 시간
	landingTime := (2 * b.VZ) / b.Gravity

	return geom.Point{
		X: b.X + b.VX*landingTime,
		Y: b.Y + b.VY*landingTime,
	}, true
}
