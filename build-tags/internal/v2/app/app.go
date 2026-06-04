// (C) Copyright 2026 GTN Group. All Rights Reserved.
// Created by Dilanka Rathnasiri on 2026-06-04

package app

import (
	"car-system/internal/engine"
	"car-system/internal/v2/infotainment"
)

func Run() {
	msg := engine.Accelerate()
	nav, err := infotainment.DisplayNavigation("Stuttgart")
	if err != nil {
		panic(err)
	}
	println(msg)
	println("navigation:", nav)
}
