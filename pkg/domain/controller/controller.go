package controller

import (
	"github.com/juanimoli/piccadilly/pkg/domain/http"
)

type Controller struct {
	Method string

	Path string

	Middleware []http.Handler

	Body http.Handler
}
