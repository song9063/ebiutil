package utils

import "github.com/hajimehoshi/ebiten/v2"

func ActualDeltaTime() float64 {
	tps := ebiten.ActualTPS()
	if tps <= 0 {
		tps = float64(ebiten.TPS())
	}
	dt := 1.0 / tps
	if dt > 0.1 {
		dt = 0.1
	}
	return dt
}
