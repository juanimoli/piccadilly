package mock

import (
	"github.com/stretchr/testify/mock"
)

type Middleware struct {
	mock.Mock
}

func (g Middleware) AbortTransactionWithError(err error) {
	_ = g.Called(err)
}

func (g Middleware) NextHandler() {
	_ = g.Called()
}

func (g Middleware) AbortTransaction() {
	_ = g.Called()
}

func (g Middleware) AbortTransactionWithStatus(code int, jsonObj interface{}) {
	_ = g.Called(code, jsonObj)
}
