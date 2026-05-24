package bmw

import (
	"errors"
	"fmt"

	"car-system/internal/v2/infotainment"
)

type bmwInfotainment struct{}

func New() infotainment.InfotainmentSystemV2 { return bmwInfotainment{} }

func (bmwInfotainment) DisplayNavigation(destination string) (string, error) {
	if destination == "" {
		return "", errors.New("destination is required")
	}
	return fmt.Sprintf("BMW iDrive v2: routing to %s", destination), nil
}
