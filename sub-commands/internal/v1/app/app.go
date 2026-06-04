package app

import (
	"car-system/internal/engine"
	"car-system/internal/v1/infotainment"
)

func Run(e engine.Engine, info infotainment.InfotainmentSystemV1) {
	msg := e.Accelerate()
	nav := info.DisplayNavigation()
	println(msg)
	println("navigation:", nav)
}
