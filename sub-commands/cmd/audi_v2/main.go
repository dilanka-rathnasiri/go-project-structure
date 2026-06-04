package main

import (
	audiengine "car-system/internal/engine/audi"
	v2app "car-system/internal/v2/app"
	audiinfo "car-system/internal/v2/infotainment/audi"
)

func main() {
	v2app.Run(audiengine.New(), audiinfo.New())
}
