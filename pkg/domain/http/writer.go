package http

type Json map[string]interface{}

type Writer interface {
	WriteJson(code int, obj interface{})

	WriteString(code int, format string, values ...interface{})

	WriteStatus(code int)
}
