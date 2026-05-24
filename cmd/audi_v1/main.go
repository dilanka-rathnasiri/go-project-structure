package main

import (
	audiengine "car-system/internal/engine/audi"
	v1app "car-system/internal/v1/app"
	audiinfo "car-system/internal/v1/infotainment/audi"
)

func main() {
	v1app.Run(audiengine.New(), audiinfo.New())
}
