package bmw

import "car-system/internal/engine"

type bmwEngine struct{}

func New() engine.Engine { return bmwEngine{} }

func (bmwEngine) Accelerate() string {
	return "BMW engine accelerating"
}
