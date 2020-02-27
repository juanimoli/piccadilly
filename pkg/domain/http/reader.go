package http

type Reader interface {
	GetUrl() string

	GetParameter(key string) string

	GetHeader(key string) string

	GetHeaders() map[string][]string

	GetFormData(key string) (string, bool)

	ReadBody(obj interface{}) error
}
