package server

import (
	"fmt"
	"github.com/juanimoli/piccadilly/pkg/presenter/binder/random"

	"github.com/juanimoli/piccadilly/pkg/domain/engine"
	"github.com/juanimoli/piccadilly/pkg/presenter/binder/health"
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
		random.CreateRandomBinder(),
	}
}

func RegisterRoutes(engine engine.ServerEngine, binders []engine.ControllerBinder) {
	for _, binder := range binders {
		binder.BindControllers(engine)
	}
}
