package health

import (
	http2 "net/http"

	"github.com/juanimoli/piccadilly/pkg/domain/controller"
	"github.com/juanimoli/piccadilly/pkg/domain/http"
)

func CreateGetController() controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/ping",
		Body:   CreateGetBody(),
	}
}

func CreateGetBody() http.Handler {
	return func(ctx *http.Context) {
		ctx.WriteString(http2.StatusOK, "pong")
	}
}
