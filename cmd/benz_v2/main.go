package main

import (
	benzengine "car-system/internal/engine/benz"
	v2app "car-system/internal/v2/app"
	benzinfo "car-system/internal/v2/infotainment/benz"
)

func main() {
	v2app.Run(benzengine.New(), benzinfo.New())
}
