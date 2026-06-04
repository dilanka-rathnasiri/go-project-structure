package main

import (
	bmwengine "car-system/internal/engine/bmw"
	v2app "car-system/internal/v2/app"
	bmwinfo "car-system/internal/v2/infotainment/bmw"
)

func main() {
	v2app.Run(bmwengine.New(), bmwinfo.New())
}
