package engine

import (
	"github.com/juanimoli/piccadilly/pkg/domain/http"
	net "net/http"
)

type ServerEngine interface {
	net.Handler

	Run(port string) error

	Shutdown() error

	GET(url string, handlers ...http.Handler)
	POST(url string, handlers ...http.Handler)
	PUT(url string, handlers ...http.Handler)
	PATCH(url string, handlers ...http.Handler)
	DELETE(url string, handlers ...http.Handler)

	Use(handlers ...http.Handler)
}
