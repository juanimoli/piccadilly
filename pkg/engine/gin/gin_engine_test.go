package gin_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/juanimoli/piccadilly/api/controller"
	http2 "github.com/juanimoli/piccadilly/api/http"
	"github.com/juanimoli/piccadilly/pkg/engine/gin"

	"github.com/stretchr/testify/assert"
)

func TestGetPortReturnsSpecificIfUsingEnvVar(t *testing.T) {
	_ = os.Setenv("PORT", "1234")
	assert.Equal(t, "1234", gin.GetPort())
	_ = os.Unsetenv("PORT")
}

func TestGetPortReturns8080(t *testing.T) {
	assert.Equal(t, "3000", gin.GetPort())
}

func TestRegisterMiddleware(t *testing.T) {
	e := gin.New()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test",
		Middleware: []http2.Handler{func(c *http2.Context) {
			c.WriteString(http.StatusOK, "Test response")
			c.AbortTransaction()
		}},
		Body: func(c *http2.Context) {
			panic("shouldn't reach here.")
		},
	})

	req, _ := http.NewRequest("GET", "/test", strings.NewReader(""))
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test response", w.Body.String())
}
