package mock

import "github.com/stretchr/testify/mock"

type ReaderMock struct {
	mock.Mock
}

func (g ReaderMock) GetParameter(key string) string {
	args := g.Called(key)
	return args.String(0)
}

func (g ReaderMock) GetFormData(key string) (string, bool) {
	args := g.Called(key)
	return args.String(0), args.Bool(1)
}

func (g ReaderMock) GetHeader(key string) string {
	args := g.Called(key)
	return args.String(0)
}

func (g ReaderMock) GetHeaders() map[string][]string {
	args := g.Called()
	return args.Get(0).(map[string][]string)
}

func (g ReaderMock) ReadBody(obj interface{}) error {
	args := g.Called(obj)
	return args.Error(0)
}

func (g ReaderMock) GetUrl() string {
	args := g.Called()
	return args.String(0)
}
