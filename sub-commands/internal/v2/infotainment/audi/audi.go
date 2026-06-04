package audi

import (
	"errors"
	"fmt"

	"car-system/internal/v2/infotainment"
)

type audiInfotainment struct{}

func New() infotainment.InfotainmentSystemV2 { return audiInfotainment{} }

func (audiInfotainment) DisplayNavigation(destination string) (string, error) {
	if destination == "" {
		return "", errors.New("destination is required")
	}
	return fmt.Sprintf("Audi MMI v2: routing to %s", destination), nil
}
