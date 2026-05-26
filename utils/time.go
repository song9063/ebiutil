package utils

import "github.com/hajimehoshi/ebiten/v2"

func ActualDeltaTime() float64 {
	tps := ebiten.ActualTPS()
	if tps <= 0 {
		tps = float64(ebiten.TPS())
	}
	return 1.0 / tps
}
