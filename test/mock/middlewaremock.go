package mock

import (
	"github.com/stretchr/testify/mock"
)

type MiddlewareMock struct {
	mock.Mock
}

func (g MiddlewareMock) AbortTransactionWithError(err error) {
	_ = g.Called(err)
}

func (g MiddlewareMock) NextHandler() {
	_ = g.Called()
}

func (g MiddlewareMock) AbortTransaction() {
	_ = g.Called()
}

func (g MiddlewareMock) AbortTransactionWithStatus(code int, jsonObj interface{}) {
	_ = g.Called(code, jsonObj)
}
