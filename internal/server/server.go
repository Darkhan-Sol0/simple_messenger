package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simple_messenger/internal/config"
	"simple_messenger/internal/datasource/database"
	"simple_messenger/internal/service"
	"simple_messenger/internal/web"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	server struct {
		cfg config.Config
		db  database.Database
	}

	Server interface {
		Run() error
	}
)

func New() Server {
	return &server{
		cfg: config.GetConfig(),
	}
}

func (s *server) start(e *echo.Echo) {
	log.Printf("Starting server at %s\n", s.cfg.GetAddress())
	if err := e.Start(s.cfg.GetAddress()); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func (s *server) shutdown(e *echo.Echo, ctx context.Context) {
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	log.Printf("Shutting down the server...")
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Printf("Graceful shutdown failed with error: %v\n", err)
	}
	log.Println("Server gracefully stopped.")
}

func (s *server) Run() (err error) {
	e := echo.New()
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	s.db, err = database.New(ctx, s.cfg)
	if err != nil {
		log.Fatalf("error connect database: %v\n", err)
	}
	defer s.db.Close(ctx)
	service := service.New(s.db)
	r := web.New(service)
	r.RegisterRoutes(e)
	go func() {
		s.start(e)
		stop()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("Context canceled, stopping server...")
	case <-quit:
		log.Println("Received termination signal, stopping server...")
	}
	s.shutdown(e, ctx)

	return nil
}
