// (C) Copyright 2026 GTN Group. All Rights Reserved.
// Created by Dilanka Rathnasiri on 2026-06-04

//go:build bmw

package infotainment

import (
	"errors"
	"fmt"
)

func DisplayNavigation(destination string) (string, error) {
	if destination == "" {
		return "", errors.New("destination is required")
	}
	return fmt.Sprintf("BMW iDrive v2: routing to %s", destination), nil
}
