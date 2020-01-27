package health

import (
	"github.com/juanimoli/piccadilly/api/engine"
	"github.com/juanimoli/piccadilly/pkg/controller/health"
)

type healthBinder struct{}

func (h *healthBinder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(health.CreateGetController())
}

func CreateHealthBinder() engine.ControllerBinder {
	return &healthBinder{}
}
