package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/juanimoli/piccadilly/api/http"
)

func CreateContext(context *gin.Context) *http.Context {
	return &http.Context{
		Reader:     CreateReader(context),
		Writer:     CreateWriter(context),
		Middleware: CreateMiddleware(context),
	}
}
