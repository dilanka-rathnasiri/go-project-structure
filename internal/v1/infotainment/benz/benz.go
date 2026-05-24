package benz

import "car-system/internal/v1/infotainment"

type benzInfotainment struct{}

func New() infotainment.InfotainmentSystemV1 { return benzInfotainment{} }

func (benzInfotainment) DisplayNavigation() string {
	return "Benz MBUX: navigation map (v1)"
}
