package gin

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/juanimoli/piccadilly/api/controller"
	"github.com/juanimoli/piccadilly/api/engine"
	gin2 "github.com/juanimoli/piccadilly/pkg/http/gin"

	"github.com/gin-gonic/gin"
)

func New() engine.ServerEngine {
	return &serverEngine{
		engine: gin.Default(),
		port:   GetPort(),
	}
}

type serverEngine struct {
	*http.Server
	engine *gin.Engine
	port   string
}

func (server serverEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.engine.ServeHTTP(writer, request)
}

func (server serverEngine) Run() error {
	if server.Server != nil {
		return errors.New("can't ignite, server already running")
	}

	addr := ":" + server.port
	fmt.Printf("Listening and serving HTTP on %s\n", addr)

	server.Server = &http.Server{
		Addr:    addr,
		Handler: server.engine,
	}

	return server.ListenAndServe()
}

func (server serverEngine) Shutdown() error {
	if server.Server == nil {
		return errors.New("no server running")
	}

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server exiting")

	return nil
}

func (server serverEngine) Register(controller controller.Controller) {
	var handlers []gin.HandlerFunc
	if controller.Middleware != nil {
		handlers = append(handlers, gin2.CreateHandlers(controller.Middleware...)...)
	}
	handlers = append(handlers, gin2.CreateHandler(controller.Body))

	server.engine.Handle(
		controller.Method,
		controller.Path,
		handlers...,
	)
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
