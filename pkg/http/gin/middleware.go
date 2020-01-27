package gin

import (
	"fmt"

	"github.com/juanimoli/piccadilly/api/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

type ginMiddleware struct {
	*gin.Context
}

func (g ginMiddleware) AbortTransactionWithError(err error) {
	var httpError http.Error
	if xerrors.As(err, &httpError) {
		if httpError.Code >= 200 && httpError.Code < 300 {
			fmt.Printf(
				"Request halted with code '%d' and message '%s' when its a successful response",
				httpError.Code,
				httpError.Reason,
			)
		} else {
			g.AbortTransactionWithStatus(httpError.Code, http.Json{
				"error": httpError.Reason,
			})
		}
	} else {
		g.AbortTransactionWithError(http.CreateInternalError())
	}
}

func (g ginMiddleware) NextHandler() {
	g.Next()
}

func (g ginMiddleware) AbortTransaction() {
	if !g.IsAborted() {
		g.Abort()
	}
}

func (g ginMiddleware) AbortTransactionWithStatus(code int, jsonObj interface{}) {
	if !g.IsAborted() {
		g.AbortWithStatusJSON(code, jsonObj)
	}
}

func CreateMiddleware(ctx *gin.Context) http.Middleware {
	return &ginMiddleware{ctx}
}
