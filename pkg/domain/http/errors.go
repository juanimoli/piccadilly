package http

import "net/http"

type Error struct {
	Code   int
	Reason string
}

func (e Error) Error() string {
	return e.Reason
}
func createHttpError(code int, message string) error {
	return Error{
		Code:   code,
		Reason: message,
	}
}

func CreateBadRequestError(message string) error {
	return createHttpError(http.StatusBadRequest, message)
}

func CreateInternalError() error {
	return createHttpError(http.StatusInternalServerError, "internal error")
}

func CreateUnauthorizedError() error {
	return createHttpError(http.StatusUnauthorized, "unauthorized")
}

func CreateNotFoundError() error {
	return createHttpError(http.StatusNotFound, "not found")
}
