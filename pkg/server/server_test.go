package server_test

import (
	"errors"
	"github.com/juanimoli/piccadilly/api/controller"
	"github.com/juanimoli/piccadilly/pkg/engine/gin"
	"github.com/juanimoli/piccadilly/pkg/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type ServerEngineMock struct{}

func (server ServerEngineMock) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
func (server ServerEngineMock) Run() error {
	return nil
}
func (server ServerEngineMock) Register(controller controller.Controller) {}
func (server ServerEngineMock) Shutdown() error {
	return nil
}

type FailingServerEngineMock struct{}

func (server FailingServerEngineMock) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
func (server FailingServerEngineMock) Run() error {
	return errors.New("woops this fails")
}
func (server FailingServerEngineMock) Register(controller controller.Controller) {}
func (server FailingServerEngineMock) Shutdown() error {
	return nil
}

func TestStartApplicationSuccess(t *testing.T) {
	err := server.StartApplication(&ServerEngineMock{})

	assert.Nil(t, err)
}

func TestStartApplicationFailureIsHandled(t *testing.T) {
	err := server.StartApplication(&FailingServerEngineMock{})

	assert.NotNil(t, err)
	assert.Equal(t, "woops this fails", err.Error())
}

func NoTestHealthGet(test *testing.T) {
	engine := gin.New()
	server.RegisterRoutes(engine, server.CreateBinders())

	w := performRequest(engine, "GET", "/ping", "")

	assert.Equal(test, 200, w.Code)
	assert.Equal(test, "pong", w.Body.String())
}
