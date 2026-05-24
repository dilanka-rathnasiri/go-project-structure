package benz

import (
	"errors"
	"fmt"

	"car-system/internal/v2/infotainment"
)

type benzInfotainment struct{}

func New() infotainment.InfotainmentSystemV2 { return benzInfotainment{} }

func (benzInfotainment) DisplayNavigation(destination string) (string, error) {
	if destination == "" {
		return "", errors.New("destination is required")
	}
	return fmt.Sprintf("Benz MBUX v2: routing to %s", destination), nil
}
