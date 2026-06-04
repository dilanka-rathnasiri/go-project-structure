// (C) Copyright 2026 GTN Group. All Rights Reserved.
// Created by Dilanka Rathnasiri on 2026-06-04

//go:build audi || !(benz || bmw)

package infotainment

func DisplayNavigation() string {
	return "Audi MMI: navigation map (v1)"
}
