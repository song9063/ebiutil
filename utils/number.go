package utils

import (
	"math"
	"math/rand/v2"

	"github.com/song9063/ebiutil/geom"
)

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func DistanceSQ(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return dx*dx + dy*dy
}

func DistanceSQByPoint(p1, p2 geom.Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return dx*dx + dy*dy
}

func RandomDegreeInRange(startAngleDeg, angleSizeDeg float64) float64 {
	randVal := RandomFloat64(0, angleSizeDeg) + startAngleDeg
	if randVal >= 360.0 {
		randVal -= 360.0
	}
	return randVal
}
func RandomRadianInRange(startAngleDeg, angleSizeDeg float64) float64 {
	return RandomDegreeInRange(startAngleDeg, angleSizeDeg) * (math.Pi / 180)
}
func PositionFromAngleRad(from geom.Point, angleRad, distance float64) geom.Point {
	return geom.Point{
		X: from.X + math.Cos(angleRad)*distance,
		Y: from.Y - math.Sin(angleRad)*distance, // ebiten좌표계는 반대로
	}
}

func VelocityFromAngle(angleRad, speed float64) (vx, vy float64) {
	return math.Cos(angleRad) * speed,
		-math.Sin(angleRad) * speed
}
