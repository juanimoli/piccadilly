package http

type Middleware interface {
	NextHandler()

	IsAborted() bool

	AbortTransaction()

	AbortTransactionWithStatus(code int, jsonObj interface{})

	AbortTransactionWithError(err error)
}
