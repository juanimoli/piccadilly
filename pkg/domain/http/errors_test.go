package http_test

import (
	"testing"

	"github.com/juanimoli/piccadilly/pkg/domain/http"
	"github.com/stretchr/testify/assert"
)

func TestCreateBadRequestError(t *testing.T) {
	err := http.CreateBadRequestError("test message")

	assert.Equal(t, 400, err.(http.Error).Code)
	assert.Equal(t, "test message", err.(http.Error).Reason)
}

func TestCreateInternalError(t *testing.T) {
	err := http.CreateInternalError()

	assert.Equal(t, 500, err.(http.Error).Code)
	assert.Equal(t, "internal error", err.(http.Error).Reason)
}
