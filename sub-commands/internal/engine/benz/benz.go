package benz

import "car-system/internal/engine"

type benzEngine struct{}

func New() engine.Engine { return benzEngine{} }

func (benzEngine) Accelerate() string {
	return "Benz engine accelerating"
}
