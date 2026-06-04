package audi

import "car-system/internal/v1/infotainment"

type audiInfotainment struct{}

func New() infotainment.InfotainmentSystemV1 { return audiInfotainment{} }

func (audiInfotainment) DisplayNavigation() string {
	return "Audi MMI: navigation map (v1)"
}
