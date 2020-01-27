package engine

import (
	"github.com/juanimoli/piccadilly/api/controller"
	"net/http"
)

type ServerEngine interface {
	ControllerRegistrable
	http.Handler

	Run() error

	Shutdown() error
}

type ControllerRegistrable interface {
	Register(controller controller.Controller)
}

type ControllerBinder interface {
	BindControllers(controllerRegistrable ControllerRegistrable)
}
