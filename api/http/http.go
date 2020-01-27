package http

type Handler func(ctx *Context)

type Context struct {
	Reader
	Writer
	Middleware
}
