package mock

import "github.com/stretchr/testify/mock"

type Reader struct {
	mock.Mock
}

func (g Reader) GetParameter(key string) string {
	args := g.Called(key)
	return args.String(0)
}

func (g Reader) GetFormData(key string) (string, bool) {
	args := g.Called(key)
	return args.String(0), args.Bool(1)
}

func (g Reader) GetHeader(key string) string {
	args := g.Called(key)
	return args.String(0)
}

func (g Reader) GetHeaders() map[string][]string {
	args := g.Called()
	return args.Get(0).(map[string][]string)
}

func (g Reader) ReadBody(obj interface{}) error {
	args := g.Called(obj)
	return args.Error(0)
}

func (g Reader) GetUrl() string {
	args := g.Called()
	return args.String(0)
}
