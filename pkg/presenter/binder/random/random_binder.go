package random

import (
	"github.com/juanimoli/piccadilly/pkg/domain/engine"
	"github.com/juanimoli/piccadilly/pkg/presenter/controller/random"
)

type randomBinder struct{}

func (h *randomBinder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(random.CreatePostController())
}

func CreateRandomBinder() engine.ControllerBinder {
	return &randomBinder{}
}

