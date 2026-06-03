package bmw

import "car-system/internal/v1/infotainment"

type bmwInfotainment struct{}

func New() infotainment.InfotainmentSystemV1 { return bmwInfotainment{} }

func (bmwInfotainment) DisplayNavigation() string {
	return "BMW iDrive: navigation map (v1)"
}
