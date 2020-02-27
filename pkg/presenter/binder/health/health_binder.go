package health

import (
	"github.com/juanimoli/piccadilly/pkg/domain/engine"
	"github.com/juanimoli/piccadilly/pkg/presenter/controller/health"
)

type healthBinder struct{}

func (h *healthBinder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(health.CreateGetController())
}

func CreateHealthBinder() engine.ControllerBinder {
	return &healthBinder{}
}
