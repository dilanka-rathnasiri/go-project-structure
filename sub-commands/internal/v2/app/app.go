package app

import (
	"car-system/internal/engine"
	"car-system/internal/v2/infotainment"
)

func Run(e engine.Engine, info infotainment.InfotainmentSystemV2) {
	msg := e.Accelerate()
	nav, err := info.DisplayNavigation("Stuttgart")
	if err != nil {
		panic(err)
	}
	println(msg)
	println("navigation:", nav)
}
