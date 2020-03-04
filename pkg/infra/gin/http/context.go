package http

import (
	"github.com/gin-gonic/gin"
	"github.com/juanimoli/piccadilly/pkg/domain/http"
)

func CreateContext(context *gin.Context) *http.Context {
	return &http.Context{
		Reader:     CreateReader(context),
		Writer:     CreateWriter(context),
		Middleware: CreateMiddleware(context),
	}
}
