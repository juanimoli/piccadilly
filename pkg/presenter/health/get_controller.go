package health

import (
	http2 "net/http"

	"github.com/juanimoli/piccadilly/pkg/domain/http"
)

func CreateGetBody() http.Handler {
	return func(ctx *http.Context) {
		ctx.WriteString(http2.StatusOK, "pong")
	}
}
