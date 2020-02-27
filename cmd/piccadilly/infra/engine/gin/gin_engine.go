package gin

import (
	"context"
	"errors"
	"fmt"
	gin2 "github.com/juanimoli/piccadilly/cmd/piccadilly/infra/http/gin"
	http2 "github.com/juanimoli/piccadilly/pkg/domain/http"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/juanimoli/piccadilly/pkg/domain/engine"

	"github.com/gin-gonic/gin"
)

func New() engine.ServerEngine {
	return &serverEngine{
		engine: gin.Default(),
	}
}

type serverEngine struct {
	*http.Server
	engine *gin.Engine
}

func (s *serverEngine) GET(url string, handlers ...http2.Handler) {
	s.engine.GET(url, gin2.CreateHandlers(handlers...)...)
}

func (s *serverEngine) POST(url string, handlers ...http2.Handler) {
	s.engine.POST(url, gin2.CreateHandlers(handlers...)...)
}

func (s *serverEngine) PUT(url string, handlers ...http2.Handler) {
	s.engine.PUT(url, gin2.CreateHandlers(handlers...)...)
}

func (s *serverEngine) PATCH(url string, handlers ...http2.Handler) {
	s.engine.PATCH(url, gin2.CreateHandlers(handlers...)...)
}

func (s *serverEngine) DELETE(url string, handlers ...http2.Handler) {
	s.engine.DELETE(url, gin2.CreateHandlers(handlers...)...)
}

func (s *serverEngine) Use(handlers ...http2.Handler) {
	s.engine.Use(gin2.CreateHandlers(handlers...)...)
}

func (s serverEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.engine.ServeHTTP(writer, request)
}

func (s serverEngine) Run(port string) error {
	if s.Server != nil {
		return errors.New("can't ignite, server already running")
	}

	addr := ":" + port
	fmt.Printf("Listening and serving HTTP on %s\n", addr)

	s.Server = &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}

	return s.ListenAndServe()
}

func (s serverEngine) Shutdown() error {
	if s.Server == nil {
		return errors.New("no server running")
	}

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server exiting")

	return nil
}
