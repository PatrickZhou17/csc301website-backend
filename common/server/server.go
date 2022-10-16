package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"shopping-cart/common/config"
	"shopping-cart/common/log"
)

type server struct {
	*http.Server
}

var TraceLogger *log.Logger

type Server interface {
	Header(key string) string
	GetHeaders() http.Header
	GetContext() context.Context
	Get(key string) (value interface{}, exists bool)
	Log() *log.Logger
}

type severKey struct{}

type CommonRequest struct {
	ctx *gin.Context
}

func NewContext(ctx context.Context, ginCtx *gin.Context) context.Context {
	return context.WithValue(ctx, severKey{}, &CommonRequest{ctx: ginCtx})
}

func FromContext(ctx context.Context) Server {
	if s, ok := ctx.Value(severKey{}).(Server); ok {
		return s
	}
	return nil
}

func (r *CommonRequest) Get(key string) (value interface{}, exists bool) {
	return r.ctx.Get(key)
}

func (r *CommonRequest) GetHeaders() http.Header {
	return r.ctx.Request.Header.Clone()
}

func (r *CommonRequest) Header(key string) string {
	return r.ctx.GetHeader(key)
}

func (r *CommonRequest) Log() *log.Logger {
	return log.GetFromGin(r.ctx)
}

func (r *CommonRequest) GetContext() context.Context {
	return r.ctx
}

func NewServer(handler http.Handler) *server {
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.AppConf.App.Port),
		Handler: handler,
	}
	if TraceLogger == nil {
		logger := log.NewLogger("log", "trace", log.InfoLevel, true)
		TraceLogger = logger
	}

	return &server{
		s,
	}
}

func (s *server) Run() {
	go func() {
		fmt.Printf("Server starting at http://127.0.0.1:%d\n", config.AppConf.App.Port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe err: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	fmt.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)

	}
	fmt.Println("Server exiting")
}
