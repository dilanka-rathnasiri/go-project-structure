package main

import (
	benzengine "car-system/internal/engine/benz"
	v1app "car-system/internal/v1/app"
	benzinfo "car-system/internal/v1/infotainment/benz"
)

func main() {
	v1app.Run(benzengine.New(), benzinfo.New())
}
