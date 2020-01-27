package server

import (
	"fmt"

	"github.com/juanimoli/piccadilly/api/engine"
	"github.com/juanimoli/piccadilly/pkg/binder/health"
)

func StartApplication(engine engine.ServerEngine) error {
	RegisterRoutes(engine, CreateBinders())

	if err := engine.Run(); err != nil {
		_ = fmt.Errorf("error running server %s", err.Error())
		return err
	}
	return nil
}

func CreateBinders() []engine.ControllerBinder {
	return []engine.ControllerBinder{
		health.CreateHealthBinder(),
	}
}

func RegisterRoutes(engine engine.ServerEngine, binders []engine.ControllerBinder) {
	for _, binder := range binders {
		binder.BindControllers(engine)
	}
}
