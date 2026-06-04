// (C) Copyright 2026 GTN Group. All Rights Reserved.
// Created by Dilanka Rathnasiri on 2026-06-04

package app

import (
	"car-system/internal/engine"
	"car-system/internal/v1/infotainment"
)

func Run() {
	msg := engine.Accelerate()
	nav := infotainment.DisplayNavigation()
	println(msg)
	println("navigation:", nav)
}
