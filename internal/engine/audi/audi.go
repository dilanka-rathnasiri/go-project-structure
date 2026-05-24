package audi

import "car-system/internal/engine"

type audiEngine struct{}

func New() engine.Engine { return audiEngine{} }

func (audiEngine) Accelerate() string {
	return "Audi engine accelerating"
}
