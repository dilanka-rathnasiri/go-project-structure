package main

import (
	bmwengine "car-system/internal/engine/bmw"
	v1app "car-system/internal/v1/app"
	bmwinfo "car-system/internal/v1/infotainment/bmw"
)

func main() {
	v1app.Run(bmwengine.New(), bmwinfo.New())
}
