package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/juanimoli/piccadilly/api/http"
)

func CreateHandler(handler http.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		handler(CreateContext(context))
	}
}

func CreateHandlers(handlers ...http.Handler) []gin.HandlerFunc {
	var ginHandlers []gin.HandlerFunc
	for _, handler := range handlers {
		ginHandlers = append(ginHandlers, CreateHandler(handler))
	}
	return ginHandlers
}
